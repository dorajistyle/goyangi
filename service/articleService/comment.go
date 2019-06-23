package articleService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/form"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/commentService"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateArticleCommentCount updates article's comment count.
func UpdateArticleCommentCount(article *model.Article) (int, error) {
	article.CommentCount = db.ORM.Model(article).Association("Comments").Count()
	if db.ORM.Save(article).Error != nil {
		return http.StatusInternalServerError, errors.New("Article comment's count is not updated.")
	}
	return http.StatusOK, nil
}

// CreateCommentOnArticle creates a comment to an article.
func CreateCommentOnArticle(c *gin.Context) (int, error) {
	article := &model.Article{}
	status, err := commentService.CreateComment(c, article)
	if err != nil {
		return status, err
	}
	status, err = UpdateArticleCommentCount(article)
	if err != nil {
		return status, err
	}
	return http.StatusCreated, nil
}

// RetrieveCommentsOnArticles retrieve comments on an article.
func RetrieveCommentsOnArticle(c *gin.Context) (model.CommentList, int, error) {
	var article model.Article
	var commentList model.CommentList
	var retrieveListForm form.RetrieveListForm

	articleId := c.Params.ByName("id")
	if db.ORM.First(&article, articleId).RecordNotFound() {
		return commentList, http.StatusNotFound, errors.New("Article is not found.")
	}
	bindErr := c.MustBindWith(&retrieveListForm, binding.Form)
	log.Debugf("[RetrieveCommentsOnArticle] bind error : %s\n", bindErr)
	if bindErr != nil {
		return commentList, http.StatusBadRequest, errors.New("Comments are not retrieved.")
	}
	log.Debugf("retrieveListForm : %v", retrieveListForm)
	commentList = commentService.RetrieveComments(article, retrieveListForm.CurrentPage)
	return commentList, http.StatusOK, nil
}

// UpdateCommentOnArticle updates a comment on an article.
func UpdateCommentOnArticle(c *gin.Context) (int, error) {
	article := &model.Article{}
	status, err := commentService.UpdateComment(c, article)
	if err != nil {
		return status, err
	}
	status, err = UpdateArticleCommentCount(article)
	if err != nil {
		return status, err
	}
	return http.StatusOK, err
}

// DeleteCommentOnArticle deletes a comment from an article.
func DeleteCommentOnArticle(c *gin.Context) (int, error) {
	article := &model.Article{}
	status, err := commentService.DeleteComment(c, article)
	if err != nil {
		return status, err
	}
	status, err = UpdateArticleCommentCount(article)
	if err != nil {
		return status, err
	}
	return http.StatusOK, err
}
