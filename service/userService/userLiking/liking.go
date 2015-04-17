package userLiking

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/form"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/likingService"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/pagination"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateUserLikingCount updates user liking count.
func UpdateUserLikingCount(user *model.User) (int, error) {
	log.Debug("UpdateUserLikingCount performed")
	user.LikingCount = db.ORM.Model(user).Association("Likings").Count()
	if db.ORM.Save(user).Error != nil {
		return http.StatusInternalServerError, errors.New("User liking count is not updated.")
	}
	return http.StatusOK, nil
}

// UpdateUserLikedCount updates user liked count.
func UpdateUserLikedCount(c *gin.Context) (int, error) {
	log.Debug("UpdateUserLikedCount performed")
	currentUserSrc, err := userService.CurrentUser(c)
	var currentUser model.User
	if err != nil {
		return http.StatusUnauthorized, err
	}
	db.ORM.First(&currentUser, currentUserSrc.Id)
	currentUser.LikedCount = db.ORM.Model(currentUser).Association("Liked").Count()
	log.Debugf("LikedCount : %d", currentUser.LikedCount)
	if db.ORM.Save(currentUser).Error != nil {
		return http.StatusInternalServerError, errors.New("User liked count is not updated.")
	}
	return http.StatusOK, nil
}

// CreateLikingOnUser creates liking on user.
func CreateLikingOnUser(c *gin.Context) (int, error) {
	user := &model.User{}
	status, err := likingService.CreateLiking(c, user)
	if err != nil {
		return status, err
	}
	status, err = UpdateUserLikingCount(user)
	if err != nil {
		return status, err
	}
	// if err == nil {
	status, err = UpdateUserLikedCount(c)
	return status, err
	// }
	// if err != nil {
	// 	return err, 400
	// }
}

// DeleteLikingOnUser deletes liking on a user.
func DeleteLikingOnUser(c *gin.Context) (int, error) {
	user := &model.User{}
	status, err := likingService.DeleteLiking(c, user)
	if err != nil {
		return status, err
	}
	// if err == nil {
	log.Debug("DeleteLikingOnUser likingDeleted")
	status, err = UpdateUserLikingCount(user)
	if err != nil {
		return status, err
	}
	status, err = UpdateUserLikedCount(c)
	// if err != nil {
	return status, err
	// }

	// 	log.Debugf("DeleteLikingOnUser error %v", err)
	// 	if err == nil {
	// 		err = UpdateUserLikedCount(c)
	// 	}
	// // }
	// return err
}

// RetrieveLikingsOnUser retrieves likings on a user.
func RetrieveLikingsOnUser(c *gin.Context) ([]model.User, int, bool, bool, int, error) {
	var user model.User
	var likings []model.User
	var retrieveListForm form.RetrieveListForm
	var hasPrev, hasNext bool
	var offset, currentPage int
	userId := c.Params.ByName("id")
	// userId, err := strconv.Atoi(c.Params.ByName("id"))
	log.Debugf("Liking params : %v", c.Params)
	c.BindWith(&retrieveListForm, binding.Form)
	log.Debugf("retrieveListForm %+v\n", retrieveListForm)
	log.Debugf("offset %+d\n", offset)
	// if hasUserId := log.CheckError(err); hasUserId {
	if db.ORM.First(&user, userId).RecordNotFound() {
		return likings, currentPage, hasPrev, hasNext, http.StatusNotFound, errors.New("User is not found.")
	}
	offset, currentPage, hasPrev, hasNext = pagination.Paginate(retrieveListForm.CurrentPage, config.LikingPerPage, user.LikingCount)
	db.ORM.Limit(config.LikingPerPage).Offset(offset).Model(&user).Association("Likings").Find(&likings)
	// for _, element := range likings {
	// 	log.Debug("User Likings: " + string(element.Id))
	// }
	// }
	return likings, currentPage, hasPrev, hasNext, http.StatusOK, nil
}

// RetrieveLikedOnUser retrieve liked on a user.
func RetrieveLikedOnUser(c *gin.Context) ([]model.User, int, bool, bool, int, error) {
	var user model.User
	var liked []model.User
	var retrieveListForm form.RetrieveListForm
	var hasPrev, hasNext bool
	var offset, currentPage int
	userId := c.Params.ByName("id")
	// userId, err := strconv.Atoi(c.Params.ByName("id"))
	log.Debugf("Liking params : %v", c.Params)
	c.BindWith(&retrieveListForm, binding.Form)
	log.Debugf("retrieveListForm %+v\n", retrieveListForm)
	log.Debugf("offset %+d\n", offset)
	// if hasUserId := log.CheckError(err); hasUserId {
	// db.ORM.First(&user, userId)
	if db.ORM.First(&user, userId).RecordNotFound() {
		return liked, currentPage, hasPrev, hasNext, http.StatusNotFound, errors.New("User is not found.")
	}
	offset, currentPage, hasPrev, hasNext = pagination.Paginate(retrieveListForm.CurrentPage, config.LikedPerPage, user.LikedCount)
	db.ORM.Limit(config.LikingPerPage).Offset(offset).Model(&user).Association("Liked").Find(&liked)
	// }
	return liked, currentPage, hasPrev, hasNext, http.StatusOK, nil
}
