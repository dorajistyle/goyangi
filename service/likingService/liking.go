package likingService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/pagination"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
)

func RetrieveLikings(item interface{}, currentUserId uint, currentPages ...int) model.LikingList {
	var users []model.User
	var currentPage int
	currentUserlikedCount := db.ORM.Model(item).Where("id =?", currentUserId).Association("Likings").Count()
	log.Debugf("Current user like count : %d", currentUserlikedCount)
	isLiked := currentUserlikedCount == 1
	if len(currentPages) > 0 {
		currentPage = currentPages[0]
	} else {
		currentPage = 1
	}
	log.Debugf("Liking Association : %v", db.ORM.Model(item).Association("Likings"))
	count := db.ORM.Model(item).Association("Likings").Count()
	offset, currentPage, hasPrev, hasNext := pagination.Paginate(currentPage, viper.GetInt("pagination.liking"), count)
	db.ORM.Limit(viper.GetInt("pagination.liking")).Order(viper.GetString("order.liking")).Offset(offset).Select(viper.GetString("publicFields.user")).Model(item).Association("Likings").Find(&users)

	return model.LikingList{Likings: users, HasPrev: hasPrev, HasNext: hasNext, Count: count, CurrentPage: currentPage, IsLiked: isLiked}
}

// RetrieveLiked retrieves liked.
func RetrieveLiked(item interface{}, currentPages ...int) model.LikedList {
	var users []model.User
	var currentPage int
	if len(currentPages) > 0 {
		currentPage = currentPages[0]
	} else {
		currentPage = 1
	}
	count := db.ORM.Model(item).Association("Liked").Count()
	offset, currentPage, hasPrev, hasNext := pagination.Paginate(currentPage, viper.GetInt("pagination.liked"), count)
	db.ORM.Limit(viper.GetInt("pagination.liked")).Order(viper.GetString("order.liked")).Offset(offset).Select(viper.GetString("publicFields.user")).Model(item).Association("Liked").Find(&users)
	//
	// var likedArr []model.PublicUser
	// for _, user := range users {
	// 	likedArr = append(likedArr, model.PublicUser{User: user})
	// }

	log.Debugf("liked : %v", users)
	return model.LikedList{Liked: users, HasPrev: hasPrev, HasNext: hasNext, Count: count, CurrentPage: currentPage}
}

// CreateLiking create a liking.
func CreateLiking(c *gin.Context, item interface{}) (int, error) {
	var form CreateLikingForm
	var likingUser model.User

	bindErr := c.MustBindWith(&form, binding.Form)
	log.Debugf("bind error : %s\n", bindErr)
	if bindErr != nil {
		return http.StatusInternalServerError, errors.New("Invalid form.")
	}

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
