package locationService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/form"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/commentService"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateLocationCommentCount updates location's comment count.
func UpdateLocationCommentCount(location *model.Location) (int, error) {
	location.CommentCount = db.ORM.Model(location).Association("Comments").Count()
	if db.ORM.Save(location).Error != nil {
		return http.StatusInternalServerError, errors.New("Location comment's count is not updated.")
	}
	return http.StatusOK, nil
}

// CreateCommentOnLocation creates comment on location.
func CreateCommentOnLocation(c *gin.Context) (int, error) {
	location := &model.Location{}
	status, err := commentService.CreateComment(c, location)
	if err != nil {
		return status, err
	}
	status, err = UpdateLocationCommentCount(location)
	if err != nil {
		return status, err
	}
	return http.StatusCreated, nil
}

// RetrieveCommentsOnLocations retrieve comments on a location.
func RetrieveCommentsOnLocation(c *gin.Context) ([]model.Comment, int, bool, bool, int, int, error) {
	var location model.Location
	var comments []model.Comment
	var retrieveListForm form.RetrieveListForm
	var hasPrev, hasNext bool
	var currentPage, count int
	locationId := c.Params.ByName("id")
	if db.ORM.First(&location, locationId).RecordNotFound() {
		return comments, currentPage, hasPrev, hasNext, count, http.StatusNotFound, errors.New("Location is not found.")
	}

	c.BindWith(&retrieveListForm, binding.Form)
	comments, currentPage, hasPrev, hasNext, count = commentService.RetrieveComments(location, retrieveListForm.CurrentPage)
	return comments, currentPage, hasPrev, hasNext, count, http.StatusOK, nil
}

// UpdateCommentOnLocation updates a comment of an location.
func UpdateCommentOnLocation(c *gin.Context) (int, error) {
	location := &model.Location{}
	status, err := commentService.UpdateComment(c, location)
	if err != nil {
		return status, err
	}
	status, err = UpdateLocationCommentCount(location)
	if err != nil {
		return status, err
	}
	return http.StatusOK, err
}

// DeleteCommentOnLocation deletes a comment from an comment.
func DeleteCommentOnLocation(c *gin.Context) (int, error) {
	location := &model.Location{}
	status, err := commentService.DeleteComment(c, location)
	if err != nil {
		return status, err
	}
	status, err = UpdateLocationCommentCount(location)
	if err != nil {
		return status, err
	}
	return http.StatusOK, err
}
