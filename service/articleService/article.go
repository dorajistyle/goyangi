package articleService

import (
	"bytes"
	"errors"
	"time"

	"net/http"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/commentService"
	"github.com/dorajistyle/goyangi/service/likingService"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/pagination"
	"github.com/dorajistyle/goyangi/util/stringHelper"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// canUserWrite check that user can write an article.
func canUserWrite(c *gin.Context, category int) bool {
	canWrite := false
	user, err := userService.CurrentUser(c)
	if err == nil {
		if category == 100 {
			if userPermission.HasAdmin(&user) {
				canWrite = true
			}
		} else {
			canWrite = true
		}
	}
	return canWrite
}

// assignRelatedUser assign related user to the article.
func assignRelatedUser(article *model.Article) {
	var tempUser model.User
	if db.ORM.Model(&article).Select(config.UserPublicFields).Related(&tempUser).RecordNotFound() {
		log.Warn("user is not found.")
	}
	// article.Author = model.PublicUser{User: &tempUser}
	article.Author = tempUser
}

// CreateArticle creates an article.
func CreateArticle(c *gin.Context) (model.Article, int, error) {
	var form ArticleForm
	bindErr := c.MustBindWith(&form, binding.Form)
	log.Debugf("bind error : %s\n", bindErr)

	user, _ := userService.CurrentUser(c)
	form.UserId = user.Id
	article := model.Article{}
	modelHelper.AssignValue(&article, &form)
	log.Debugf("article : %s\n", article)
	if bindErr != nil || db.ORM.Create(&article).Error != nil {
		return article, http.StatusInternalServerError, errors.New("Article is not created.")
	}
	return article, http.StatusCreated, nil
}

// CreateArticles creates articles.
func CreateArticles(c *gin.Context) (int, error) {
	var forms ArticlesForm
	c.BindJSON(&forms)
	log.Debugf("CreateFiles c form : %v", forms)

	user, _ := userService.CurrentUser(c)
	sqlStrBuffer := new(bytes.Buffer)
	stringHelper.Concat(sqlStrBuffer, "INSERT INTO article(user_id, title, url, content, image_name, category_id, created_at) VALUES ")
	values := []interface{}{}
	for _, article := range forms.Articles {
		stringHelper.Concat(sqlStrBuffer, "(?, ?, ?, ?, ?, ?, ?),")
		values = append(values, user.Id, article.Title, article.Url, article.Content, article.ImageName, 0, time.Now())
	}
	// sqlStrBuffer.Truncate(sqlStrBuffer.Len() - 1) is slower than slice.
	if len(values) > 0 {
		sqlStr := sqlStrBuffer.String()
		sqlStr = sqlStr[0 : len(sqlStr)-1]
		log.Debugf("sqlStr for File : %s", sqlStr)
		db.ORM.Exec(sqlStr, values...)
	}

	return http.StatusCreated, nil
}

// RetrieveArticle retrieve an article.
func RetrieveArticle(c *gin.Context) (model.Article, int, error) {
	var article model.Article
	var count int
	id := c.Params.ByName("id")
	if db.ORM.First(&article, "id = ?", id).RecordNotFound() {
		return article, http.StatusNotFound, errors.New("Article is not found.")
	}
	log.Debugf("Article : %s\n", article)
	log.Debugf("Count : %s\n", count)
	currentUser, _ := userService.CurrentUser(c)

	assignRelatedUser(&article)

	article.CommentList = commentService.RetrieveComments(article)
	var likingList model.LikingList
	likingList = likingService.RetrieveLikings(article, currentUser.Id)

	// DEPRECATED likingMeta.SetLikingPageMeta(&likingList, currentPage, hasPrev, hasNext, article.LikingCount, currentUserlikedCount)
	article.LikingList = likingList
	return article, http.StatusOK, nil
}

// RetrieveArticles retrieves articles.
func RetrieveArticles(c *gin.Context) (model.ArticleList, int, error) {
	var articles []model.Article
	var category int
	var articleCount, articlePerPage int

	filter := &ArticleFilter{}

	articlePerPage = config.ArticlePerPage

	whereBuffer := new(bytes.Buffer)
	whereValues := []interface{}{}
	if c.Bind(filter) == nil {
		log.Debug("====== Bind By Query String ======")
		if filter.UserId > 0 {
			stringHelper.Concat(whereBuffer, "user_id = ?")
			whereValues = append(whereValues, filter.UserId)
			log.Debugf("userId : %d\n", filter.UserId)
		}
		if len(filter.Categories) > 0 {
			if len(whereValues) == 1 {
				stringHelper.Concat(whereBuffer, " and ")
			}
			stringHelper.Concat(whereBuffer, "category_id = ?")
			whereValues = append(whereValues, filter.Categories[0])
			log.Debugf("categories : %d\n", filter.Categories[0])
			category = filter.Categories[0]
		}
		if filter.ArticlePerPage > 0 {
			articlePerPage = filter.ArticlePerPage
			log.Debugf("articlePerPage : %d\n", filter.ArticlePerPage)
		}
	} else {
		log.Debug("no filters found.\n")
	}

	log.Debugf("filter %v.\n", filter)
	whereStr := whereBuffer.String()
	log.Debugf("whereStr %s.\n", whereStr)
	log.Debugf("whereValues %v.\n", whereValues)
	db.ORM.Model(model.Article{}).Where(whereStr, whereValues...).Count(&articleCount)
	offset, currentPage, hasPrev, hasNext := pagination.Paginate(filter.CurrentPage, articlePerPage, articleCount)
	log.Debugf("currentPage, perPage, total : %d, %d, %d", filter.CurrentPage, articlePerPage, articleCount)
	log.Debugf("offset, currentPage, hasPrev, hasNext : %d, %d, %t, %t", offset, currentPage, hasPrev, hasNext)
	db.ORM.Limit(articlePerPage).Offset(offset).Order(config.ArticleOrder).Where(whereStr, whereValues...).Find(&articles)
	return model.ArticleList{Articles: articles, Category: category, HasPrev: hasPrev, HasNext: hasNext, Count: articleCount, CurrentPage: currentPage, PerPage: articlePerPage}, http.StatusOK, nil
}

// UpdateArticle updates an article.
func UpdateArticle(c *gin.Context) (model.Article, int, error) {
	var article model.Article
	var form ArticleForm
	bindErr := c.MustBindWith(&form, binding.Form)
	log.Debugf("[Update] bind error : %s\n", bindErr)
	if bindErr != nil {
		return article, http.StatusBadRequest, errors.New("Article is not updated.")
	}
	if db.ORM.First(&article, form.Id).RecordNotFound() {
		log.Debug("[Update] Record Not found\n")
		return article, http.StatusNotFound, errors.New("Article is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, article.UserId)
	if err != nil {
		log.Debugf("[Update] User err : %s\n" + err.Error())
		return article, status, err
	}
	article.Title = form.Title
	article.Url = form.Url
	article.ImageName = form.ImageName
	article.Content = form.Content
	if db.ORM.Save(&article).Error != nil {
		log.Debug("[Update] Update save error\n")
		return article, http.StatusBadRequest, errors.New("Article is not updated.")
	}
	return article, http.StatusOK, nil
}

// DeleteArticle deletes an article.
func DeleteArticle(c *gin.Context) (int, error) {
	var article model.Article
	id := c.Params.ByName("id")
	if db.ORM.First(&article, id).RecordNotFound() {
		return http.StatusNotFound, errors.New("Article is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, article.UserId)
	if err != nil {
		return status, err
	}
	if db.ORM.Delete(&article).Error != nil {
		return http.StatusBadRequest, errors.New("Article is not deleted.")
	}
	return http.StatusOK, nil
}
