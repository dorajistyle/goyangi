package userLiking

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/form"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/pagination"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateUserLikingCount updates user liking count.
func UpdateUserLikingCount(user *model.User) (int, error) {
	log.Debug("UpdateUserLikingCount performed")
	user.LikingCount = len(user.Likings)
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
	currentUser.LikedCount = len(currentUser.Liked)
	log.Debugf("LikedCount : %d", currentUser.LikedCount)
	if db.ORM.Save(currentUser).Error != nil {
		return http.StatusInternalServerError, errors.New("User liked count is not updated.")
	}
	return http.StatusOK, nil
}

// CreateLikingForm is used when creating a liking.
type CreateLikingForm struct {
	UserId   uint `form:"userId" binding:"required"`
	ParentId uint `form:"parentId" binding:"required"`
}

// CreateLiking create a liking to an user.
func CreateLiking(c *gin.Context, user interface{}) (int, error) {
	var form CreateLikingForm
	var likingUser model.User
	c.BindWith(&form, binding.Form)
	log.Debugf("liking_form : %v", form)
	if db.ORM.First(user, form.ParentId).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found.")
	}
	if db.ORM.First(&likingUser, form.UserId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Follower is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, likingUser.Id)
	if err != nil {
		return status, err
	}
	var usersFollowers = model.UsersFollowers{UserID: form.ParentId, FollowerID: form.UserId}
	var likingUserCount int
	db.ORM.Model(&model.UsersFollowers{}).Where("user_id = ? and follower_id = ?", form.ParentId, form.UserId).Count(&likingUserCount)
	if likingUserCount != 0 {
		return http.StatusInternalServerError, errors.New("User already followed.")
	}
	if db.ORM.Create(&usersFollowers).Error != nil {
		return http.StatusBadRequest, nil
	}
	return http.StatusOK, nil
}

// DeleteLiking deletes liking of an user.
func DeleteLiking(c *gin.Context, inputUser interface{}) (int, error) {
	userId := c.Params.ByName("id")
	followerId := c.Params.ByName("userId")
	var user model.User
	var likingUser model.User
	log.Debugf("user id : %d , follower id : %d \n", userId, followerId)
	if db.ORM.First(&user, userId).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found.")
	}
	if db.ORM.First(&likingUser, followerId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Follower is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, likingUser.Id)
	if err != nil {
		return status, err
	}
	var usersFollowers = model.UsersFollowers{UserID: user.Id, FollowerID: likingUser.Id}
	db.ORM.Delete(&usersFollowers)
	var likingUserCount int
	db.ORM.Model(&model.UsersFollowers{}).Where("user_id = ? and follower_id = ?", user.Id, likingUser.Id).Count(&likingUserCount)
	if likingUserCount != 0 {
		return http.StatusInternalServerError, errors.New("User following is not deleted.")
	}
	return http.StatusOK, nil
}

// CreateLikingOnUser creates liking on user.
func CreateLikingOnUser(c *gin.Context) (int, error) {
	user := &model.User{}
	status, err := CreateLiking(c, user)
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
	status, err := DeleteLiking(c, user)
	if err != nil {
		return status, err
	}
	log.Debug("DeleteLikingOnUser likingDeleted")
	status, err = UpdateUserLikingCount(user)
	if err != nil {
		return status, err
	}
	status, err = UpdateUserLikedCount(c)
	return status, err

}

// RetrieveLikingsOnUser retrieves likings on a user.
func RetrieveLikingsOnUser(c *gin.Context) ([]model.User, int, bool, bool, int, error) {
	var user model.User
	var likings []model.User
	var retrieveListForm form.RetrieveListForm
	var hasPrev, hasNext bool
	var offset, currentPage int
	userId := c.Params.ByName("id")
	log.Debugf("Liking params : %v", c.Params)
	c.BindWith(&retrieveListForm, binding.Form)
	log.Debugf("retrieveListForm %+v\n", retrieveListForm)
	log.Debugf("offset %+d\n", offset)
	// if hasUserId := log.CheckError(err); hasUserId {
	if db.ORM.First(&user, userId).RecordNotFound() {
		return likings, currentPage, hasPrev, hasNext, http.StatusNotFound, errors.New("User is not found.")
	}
	offset, currentPage, hasPrev, hasNext = pagination.Paginate(retrieveListForm.CurrentPage, config.LikingPerPage, user.LikingCount)
	db.ORM.Limit(config.LikingPerPage).Offset(offset).
		Joins("JOIN users_followers on users_followers.user_id=?", user.Id).
		Where("users.id = users_followers.follower_id").
		Group("users.id").Find(&likings)
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
	log.Debugf("Liked params : %v", c.Params)
	c.BindWith(&retrieveListForm, binding.Form)
	log.Debugf("retrieveListForm %+v\n", retrieveListForm)
	log.Debugf("offset %+d\n", offset)

	if db.ORM.First(&user, userId).RecordNotFound() {
		return liked, currentPage, hasPrev, hasNext, http.StatusNotFound, errors.New("User is not found.")
	}
	offset, currentPage, hasPrev, hasNext = pagination.Paginate(retrieveListForm.CurrentPage, config.LikedPerPage, user.LikedCount)
	db.ORM.Limit(config.LikedPerPage).Offset(offset).
		Joins("JOIN users_followers on users_followers.follower_id=?", user.Id).
		Where("users.id = users_followers.user_id").
		Group("users.id").Find(&liked)

	return liked, currentPage, hasPrev, hasNext, http.StatusOK, nil
}
