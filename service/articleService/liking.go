package articleService

import (
	"errors"
	"net/http"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/form"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/service/likingService"
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
func RetrieveLikingsOnArticles(c *gin.Context) (model.LikingList, int, error) {
	var article model.Article
	var likingList model.LikingList
	var retrieveListForm form.RetrieveListForm
	
	articleId := c.Params.ByName("id")
	log.Debugf("Liking params : %v", c.Params)

	bindErr := c.MustBindWith(&retrieveListForm, binding.Form)
	log.Debugf("[RetrieveLikingsOnArticles] bind error : %s\n", bindErr)
	if bindErr != nil {
		return likingList, http.StatusBadRequest, errors.New("Comments are not retrieved.")
	}


	log.Debugf("retrieveListForm %+v\n", retrieveListForm)
	if db.ORM.First(&article, articleId).RecordNotFound() {
		return likingList, http.StatusNotFound, errors.New("Article is not found.")
	}
	currentUser, _ := userService.CurrentUser(c)
	likingList = likingService.RetrieveLikings(article, currentUser.Id)
	// DEPRECATED likingMeta.SetLikingPageMeta(&likingList, currentPage, hasPrev, hasNext, article.LikingCount, currentUserlikedCount)

	return likingList, http.StatusOK, nil
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
