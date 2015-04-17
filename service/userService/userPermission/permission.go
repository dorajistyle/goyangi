package userPermission

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/api/response"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
)

// HasAdmin checks that user has an admin permission.
func HasAdmin(user *model.User) bool {
	name := "admin"
	for _, role := range user.Roles {
		log.Debugf("HasAdmin role.Name : %s", role.Name)
		if role.Name == name {
			return true
		}
	}
	return false
}

// CurrentOrAdmin check that user has admin permission or user is the current user.
func CurrentOrAdmin(user *model.User, userId int64) bool {
	log.Debugf("user.Id == userId %d %d %s", user.Id, userId, user.Id == userId)
	return (HasAdmin(user) || user.Id == userId)
}

// CurrentUserIdentical check that userId is same as current user's Id.
func CurrentUserIdentical(c *gin.Context, userId int64) (int, error) {
	currentUser, err := userService.CurrentUser(c)
	if err != nil {
		return http.StatusUnauthorized, errors.New("Auth failed.")
	}
	if currentUser.Id != userId {
		return http.StatusForbidden, errors.New("User is not identical.")
	}

	return http.StatusOK, nil
}

// AuthRequired run function when user logged in.
func AuthRequired(f func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := userService.CurrentUser(c)
		if err != nil {
			log.Error("Auth failed.")
			response.KnownErrorJSON(c, http.StatusUnauthorized, "error.loginPlease", errors.New("Auth failed."))
			return
		}
		f(c)
		return
	}
}

// AdminRequired run function when user logged in and user has an admin role.
func AdminRequired(f func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := userService.CurrentUser(c)
		if err == nil {
			if HasAdmin(&user) {
				f(c)
				log.Debug("User has admin role.")
				return
			}
		}
		log.Error("Admin role required.")
		response.KnownErrorJSON(c, http.StatusUnauthorized, "error.adminRequired", errors.New("Admin role required."))
		return
	}
}
