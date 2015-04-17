package commentService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/pagination"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// AssignRelatedUser assign related user of comment.
func AssignRelatedUser(comments []model.Comment) {
	for i, comment := range comments {
		var tempUser model.User
		if db.ORM.Model(&comment).Select(config.UserPublicFields).Related(&tempUser).RecordNotFound() {
			log.Warn("user is not found.")
		}
		comments[i].User = model.PublicUser{User: &tempUser}
	}
}

// SetCommentPageMeta set comment's page meta.
func SetCommentPageMeta(commentList *model.CommentList, currentPage int, hasPrev bool, hasNext bool, count int) {
	if len(commentList.Comments) == 0 {
		commentList.Comments = make([]model.Comment, 0)
	}
	commentList.CurrentPage = currentPage
	commentList.HasPrev = hasPrev
	commentList.HasNext = hasNext
	commentList.Count = count
}

// CreateComment creates a comment.
func CreateComment(c *gin.Context, item interface{}) (int, error) {
	var form CreateCommentForm
	var comment model.Comment
	c.BindWith(&form, binding.Form)
	log.Debugf("comment_form : %v", form)
	status, err := userPermission.CurrentUserIdentical(c, form.UserId)
	if err != nil {
		return status, err
	}
	if db.ORM.First(item, form.ParentId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Item is not found.")
	}
	modelHelper.AssignValue(&comment, &form)
	db.ORM.Model(item).Association("Comments").Append(comment)
	return http.StatusOK, nil
}

// RetrieveComments retrieves comments.
func RetrieveComments(item interface{}, currentPages ...int) ([]model.Comment, int, bool, bool, int) {
	var comments []model.Comment
	var currentPage int
	if len(currentPages) > 0 {
		currentPage = currentPages[0]
	} else {
		currentPage = 1
	}
	count := db.ORM.Model(item).Association("Comments").Count()
	offset, currentPage, hasPrev, hasNext := pagination.Paginate(currentPage, config.CommentPerPage, count)
	db.ORM.Limit(config.CommentPerPage).Order(config.CommentOrder).Offset(offset).Model(item).Association("Comments").Find(&comments)
	AssignRelatedUser(comments)
	log.Debugf("comments : %v", comments)
	return comments, currentPage, hasPrev, hasNext, count
}

// UpdateComment updates a comment.
func UpdateComment(c *gin.Context, item interface{}) (int, error) {
	var form CommentForm
	var comment model.Comment
	c.BindWith(&form, binding.Form)
	if db.ORM.First(item, form.ParentId).RecordNotFound() {
		return http.StatusNotFound, errors.New("Item is not found.")
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
