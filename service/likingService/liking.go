package likingService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateLiking create a liking.
func CreateLiking(c *gin.Context, item interface{}) (int, error) {
	var form CreateLikingForm
	var likingUser model.User
	c.BindWith(&form, binding.Form)
	log.Debugf("liking_form : %v", form)
	if db.ORM.First(item, form.ParentId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Item is not found.")
	}
	if db.ORM.First(&likingUser, form.UserId).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, likingUser.Id)
	if err != nil {
		return status, err
	}
	db.ORM.Model(item).Association("Likings").Append(likingUser)
	return http.StatusOK, nil
}

// DeleteLiking deletes liking.
func DeleteLiking(c *gin.Context, item interface{}) (int, error) {
	itemId := c.Params.ByName("id")
	userId := c.Params.ByName("userId")
	var likingUser model.User
	log.Debugf("item id : %d , user id : %d \n", itemId, userId)
	if db.ORM.First(item, itemId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Item is not found.")
	}
	if db.ORM.First(&likingUser, userId).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, likingUser.Id)
	if err != nil {
		return status, err
	}
	db.ORM.Model(item).Association("Likings").Delete(likingUser)
	likingUserCount := db.ORM.Where("id = ?", likingUser.Id).Model(item).Association("Likings").Count()
	if likingUserCount != 0 {
		return http.StatusInternalServerError, errors.New("Article liking's is not deleted.")
	}
	return http.StatusOK, nil
}
