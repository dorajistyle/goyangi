package commentService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/pagination"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/viper"
)

// AssignRelatedUser assign related user of comment.
func AssignRelatedUser(comments []model.Comment) {
	for i, comment := range comments {
		var tempUser model.User
		if db.ORM.Model(&comment).Select(viper.GetString("publicFields.user")).Related(&tempUser).RecordNotFound() {
			log.Warn("user is not found.")
		}
		// comments[i].User = model.PublicUser{User: &tempUser}
		comments[i].User = tempUser
	}
}

// CreateComment creates a comment.
func CreateComment(c *gin.Context, item interface{}) (int, error) {
	var form CreateCommentForm
	var comment model.Comment
	user, userErr := userService.CurrentUser(c)
	log.Debugf("user error : %s\n", userErr)
	if userErr != nil {
		return http.StatusForbidden, userErr
	}
	form.UserId = user.Id
	if db.ORM.First(item, c.Param("id")).RecordNotFound() {
		return http.StatusNotFound, errors.New("Item is not found.")
	}

	bindErr := c.MustBindWith(&form, binding.Form)
	log.Debugf("bind error : %s\n", bindErr)
	if bindErr != nil {
		return http.StatusInternalServerError, errors.New("Invalid form.")
	}

	log.Debugf("comment_form : %v", form)

	modelHelper.AssignValue(&comment, &form)
	if db.ORM.Model(item).Association("Comments").Append(comment).Error != nil {
		return http.StatusInternalServerError, errors.New("Comment is not created.")
	}
	return http.StatusOK, nil
}

// RetrieveComments retrieves comments.
func RetrieveComments(item interface{}, currentPages ...int) model.CommentList {
	var comments []model.Comment
	var currentPage int
	if len(currentPages) > 0 {
		currentPage = currentPages[0]
	} else {
		currentPage = 1
	}
	count := db.ORM.Model(item).Association("Comments").Count()
	offset, currentPage, hasPrev, hasNext := pagination.Paginate(currentPage, viper.GetInt("pagination.comment"), count)
	db.ORM.Limit(viper.GetInt("pagination.comment")).Order(viper.GetString("order.cmment")).Offset(offset).Model(item).Association("Comments").Find(&comments)
	AssignRelatedUser(comments)
	log.Debugf("comments : %v", comments)
	return model.CommentList{Comments: comments, HasPrev: hasPrev, HasNext: hasNext, Count: count, CurrentPage: currentPage}
}

// UpdateComment updates a comment.
func UpdateComment(c *gin.Context, item interface{}) (int, error) {
	var form CommentForm
	var comment model.Comment
	itemId := c.Params.ByName("id")

	if db.ORM.First(item, itemId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Item is not found.")
	}

	bindErr := c.MustBindWith(&form, binding.Form)
	log.Debugf("bind error : %s\n", bindErr)
	if bindErr != nil {
		return http.StatusInternalServerError, errors.New("Invalid form.")
	}

	if db.ORM.First(&comment, form.CommentId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Comment is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, comment.UserId)
	if err != nil {
		return status, err
	}
	comment.Content = form.Content
	if db.ORM.Save(&comment).Error != nil {
		return http.StatusInternalServerError, errors.New("Article comment's count is not updated.")
	}
	return http.StatusOK, nil
}

// DeleteComment deletes a comment.
func DeleteComment(c *gin.Context, item interface{}) (int, error) {
	itemId := c.Params.ByName("id")
	commentId := c.Params.ByName("commentId")
	var comment model.Comment
	log.Debugf("item id : %d , comment id : %d \n", itemId, commentId)
	if db.ORM.First(item, itemId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Item is not found.")
	}
	if db.ORM.First(&comment, commentId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Comment is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, comment.UserId)
	if err != nil {
		return status, err
	}
	db.ORM.Model(item).Association("Comments").Delete(comment)
	if db.ORM.Delete(&comment).Error != nil {
		return http.StatusInternalServerError, errors.New("Comment is not deleted.")
	}
	return http.StatusOK, nil
}
