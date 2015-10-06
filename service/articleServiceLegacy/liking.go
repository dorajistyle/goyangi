package articleService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/form"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/likingService"
	"github.com/dorajistyle/goyangi/service/likingService/likingRetriever"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateArticleLikingCount updates a liking count on article.
func UpdateArticleLikingCount(article *model.Article) (int, error) {
	article.LikingCount = db.ORM.Model(article).Association("Likings").Count()
	if db.ORM.Save(article).Error != nil {
		return http.StatusInternalServerError, errors.New("Article liking's count is not updated.")
	}
	return http.StatusOK, nil
}

// CreateLikingOnArticle creates a liking on article.
func CreateLikingOnArticle(c *gin.Context) (int, error) {
	article := &model.Article{}
	status, err := likingService.CreateLiking(c, article)
	if err != nil {
		return status, err
	}
	status, err = UpdateArticleLikingCount(article)
	if err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

// RetrieveLikingsOnArticles retrieves likings on article.
func RetrieveLikingsOnArticles(c *gin.Context) ([]*model.PublicUser, int, bool, bool, int, int, error) {
	var article model.Article
	var likings []*model.PublicUser
	var retrieveListForm form.RetrieveListForm
	var hasPrev, hasNext bool
	var currentPage, count int
	articleId := c.Params.ByName("id")
	log.Debugf("Liking params : %v", c.Params)
	c.BindWith(&retrieveListForm, binding.Form)
	log.Debugf("retrieveListForm %+v\n", retrieveListForm)
	if db.ORM.First(&article, articleId).RecordNotFound() {
		return likings, currentPage, hasPrev, hasNext, count, http.StatusNotFound, errors.New("Article is not found.")
	}
	likings, currentPage, hasPrev, hasNext, count = likingRetriever.RetrieveLikings(article, retrieveListForm.CurrentPage)
	return likings, currentPage, hasPrev, hasNext, count, http.StatusOK, nil
}

// DeleteLikingOnArticle deletes liking on article.
func DeleteLikingOnArticle(c *gin.Context) (int, error) {
	article := &model.Article{}
	status, err := likingService.DeleteLiking(c, article)
	if err != nil {
		return status, err
	}
	status, err = UpdateArticleLikingCount(article)
	if err != nil {
		return status, err
	}
	return http.StatusOK, nil
}
