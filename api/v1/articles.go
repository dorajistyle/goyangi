package v1

import (
	"github.com/gin-gonic/gin"
	//	"reflect"

	"github.com/dorajistyle/goyangi/api/response"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/articleService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
)

var article model.Article

// @Title Articles
// @Description Article's router group.
func Articles(parentRoute *gin.RouterGroup) {

	route := parentRoute.Group("/articles")
	route.POST("/", userPermission.AuthRequired(createArticle))
	route.POST("/all", userPermission.AuthRequired(createArticles))
	route.GET("/:id", retrieveArticle)
	route.GET("/", retrieveArticles)
	route.PUT("/:id", userPermission.AuthRequired(updateArticle))
	route.DELETE("/:id", userPermission.AuthRequired(deleteArticle))

	route.POST("/comments", userPermission.AuthRequired(createCommentOnArticle))
	route.GET("/:id/comments", retrieveCommentsOnArticle)
	route.PUT("/:id/comments/:commentId", userPermission.AuthRequired(updateCommentOnArticle))
	route.DELETE("/:id/comments/:commentId", userPermission.AuthRequired(deleteCommentOnArticle))

	route.POST("/likings", userPermission.AuthRequired(createLikingOnArticle))
	route.GET("/:id/likings", retrieveLikingsOnArticles)
	route.DELETE("/:id/likings/:userId", userPermission.AuthRequired(deleteLikingOnArticle))
}

// @Title createArticle
// @Description Create an article.
// @Accept  json
// @Param   title        form   string     true        "Article title."
// @Param   url        form   string     true        "Article url."
// @Param   content        form   string  false        "Article content."
// @Success 201 {object} model.Article "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Article is not created"
// @Resource /articles
// @Router /articles [post]
func createArticle(c *gin.Context) {
	article, status, err := articleService.CreateArticle(c)
	if err == nil {
		c.JSON(status, gin.H{"article": article})
	} else {
		messageTypes := &response.MessageTypes{InternalServerError: "article.error.notCreated"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title createArticles
// @Description Create a file.
// @Accept  json
// @Param   title        form   string     true        "Article title."
// @Param   url        form   string     true        "Article url."
// @Param   imageName        form   string     true        "Article imagename."
// @Param   content        form   string  false        "Article content."
// @Success 201 {object} model.Article "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Article is not created"
// @Resource /upload/files
// @Router /upload [post]
func createArticles(c *gin.Context) {
	status, err := articleService.CreateArticles(c)

	messageTypes := &response.MessageTypes{
		OK:                  "article.create.done",
		Unauthorized:        "article.error.unauthorized",
		InternalServerError: "article.error.notCreated",
	}
	messages := &response.Messages{OK: "Metadata of articles are created successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveArticle
// @Description Retrieve an article.
// @Accept  json
// @Param   id        path    int     true        "Article Id"
// @Success 200 {object} model.Article "OK"
// @Failure 404 {object} response.BasicResponse "Article is not found"
// @Resource /articles
// @Router /articles/{id} [get]
func retrieveArticle(c *gin.Context) {
	article, isAuthor, currentUserId, status, err := articleService.RetrieveArticle(c)
	if err == nil {
		c.JSON(status, gin.H{"article": article, "isAuthor": isAuthor, "currentUserId": currentUserId})
	} else {
		messageTypes := &response.MessageTypes{
			NotFound: "article.error.notFound",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title retrieveArticles
// @Description Retrieve article array.
// @Accept  json
// @Success 200 {array} model.Article "OK"
// @Resource /articles
// @Router /articles [get]
func retrieveArticles(c *gin.Context) {
	articles, canWrite, category, currentPage, hasPrev, hasNext, status, err := articleService.RetrieveArticles(c)
	if err == nil {
		c.JSON(status, gin.H{"articles": articles, "canWrite": canWrite, "category": category, "currentPage": currentPage,
			"hasPrev": hasPrev, "hasNext": hasNext})
	} else {
		messageTypes := &response.MessageTypes{}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title updateArticle
// @Description Update an article.
// @Accept  json
// @Param   id        path    int     true        "Article Id"
// @Success 200 {object} model.Article "OK"
// @Failure 400 {object} response.BasicResponse "Article is not updated"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Article is not found"
// @Resource /articles
// @Router /articles/{id} [put]
func updateArticle(c *gin.Context) {
	article, status, err := articleService.UpdateArticle(c)
	if err == nil {
		c.JSON(status, gin.H{"article": article})
	} else {
		messageTypes := &response.MessageTypes{
			BadRequest:   "article.view.updated.fail",
			Unauthorized: "article.error.isNotAuthor",
			NotFound:     "article.error.notFound"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title deleteArticle
// @Description Delete an article.
// @Accept  json
// @Param   id        path    int     true        "Article Id"
// @Success 200 {object} response.BasicResponse "Article deleted"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /articles
// @Router /articles/{id} [delete]
func deleteArticle(c *gin.Context) {
	status, err := articleService.DeleteArticle(c)
	messageTypes := &response.MessageTypes{
		OK:           "article.view.deleted.done",
		BadRequest:   "article.view.deleted.fail",
		Unauthorized: "article.error.isNotAuthor",
		NotFound:     "article.error.notFound"}
	messages := &response.Messages{OK: "Article is deleted successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title createCommentOnArticle
// @Description Create a comment on an article.
// @Accept  json
// @Param   articleId        form   int     true        "Article Id."
// @Param   content        form   string     true        "Comment content."
// @Param   imageName        form   string     true        "Article image name."
// @Param   description        form   string  false        "Article description."
// @Success 201 {object} response.BasicResponse "Comment created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Article is not found"
// @Resource /articles
// @Router /articles/comments [post]
func createCommentOnArticle(c *gin.Context) {
	status, err := articleService.CreateCommentOnArticle(c)
	messageTypes := &response.MessageTypes{
		Created:      "comment.created.done",
		Unauthorized: "comment.error.unauthorized",
		Forbidden:    "comment.error.forbidden",
		NotFound:     "comment.error.notFound"}
	messages := &response.Messages{OK: "Comment is created successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveCommentsOnArticle
// @Description Retrieve comments on an article.
// @Accept  json
// @Param   articleId        path    int     true        "Article Id"
// @Success 200 {array} model.Comment "Retrieve comments successfully"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /articles
// @Router /articles/{id}/comments [get]
func retrieveCommentsOnArticle(c *gin.Context) {
	comments, currentPage, hasPrev, hasNext, count, status, err := articleService.RetrieveCommentsOnArticle(c)

	if err == nil {
		c.JSON(status, gin.H{"comments": comments, "currentPage": currentPage, "hasPrev": hasPrev, "hasNext": hasNext, "count": count})
	} else {
		messageTypes := &response.MessageTypes{
			NotFound: "comment.error.notFound"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title updateCommentOnArticle
// @Description Update a comment on article.
// @Accept  json
// @Param   articleId        path    int     true        "Article Id"
// @Param   id                path    int     true        "Comment Id"
// @Success 200 {object} model.Comment "Comment updated successfully"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Comment is not updated"
// @Resource /articles
// @Router /articles/{id}/comments/{commentId} [put]
func updateCommentOnArticle(c *gin.Context) {
	status, err := articleService.UpdateCommentOnArticle(c)
	messageTypes := &response.MessageTypes{
		OK:                  "comment.updated.done",
		Unauthorized:        "comment.error.unauthorized",
		Forbidden:           "comment.error.forbidden",
		NotFound:            "comment.error.notFound",
		InternalServerError: "comment.updated.fail"}
	messages := &response.Messages{OK: "Comment is created successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title deleteCommentOnArticle
// @Description Delete a comment on article.
// @Accept  json
// @Param   articleId        path    int     true        "Article Id"
// @Param   id                path    int     true        "Comment Id"
// @Success 200 {object} response.BasicResponse "Comment is deleted successfully"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Comment is not deleted"
// @Resource /articles
// @Router /articles/{id}/comments/{commentId} [delete]
func deleteCommentOnArticle(c *gin.Context) {
	status, err := articleService.DeleteCommentOnArticle(c)
	messageTypes := &response.MessageTypes{
		OK:                  "comment.deleted.done",
		Unauthorized:        "comment.error.unauthorized",
		Forbidden:           "comment.error.forbidden",
		NotFound:            "comment.error.notFound",
		InternalServerError: "comment.deleted.fail"}
	messages := &response.Messages{OK: "Comment is deleted successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title createLikingOnArticle
// @Description Create a liking on an article.
// @Accept  json
// @Param   articleId        form   int     true        "Article Id."
// @Param   content        form   string     true        "Liking content."
// @Param   imageName        form   string     true        "Article image name."
// @Param   description        form   string  false        "Article description."
// @Success 201 {object} response.BasicResponse "Liking created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Article is not found"
// @Resource /articles
// @Router /articles/likings [post]
func createLikingOnArticle(c *gin.Context) {
	status, err := articleService.CreateLikingOnArticle(c)
	messageTypes := &response.MessageTypes{
		OK:           "liking.like.done",
		BadRequest:   "liking.like.fail",
		Unauthorized: "liking.error.unauthorized",
		NotFound:     "liking.error.notFound"}
	messages := &response.Messages{OK: "Article liking is created successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveLikingsOnArticles
// @Description Retrieve likings on an article.
// @Accept  json
// @Param   articleId        path    int     true        "Article Id"
// @Success 200 {array} model.PublicUser "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /articles
// @Router /articles/{id}/likings [get]
func retrieveLikingsOnArticles(c *gin.Context) {
	likings, currentPage, hasPrev, hasNext, count, status, err := articleService.RetrieveLikingsOnArticles(c)
	if err == nil {
		c.JSON(status, gin.H{"likings": likings, "currentPage": currentPage,
			"hasPrev": hasPrev, "hasNext": hasNext, "count": count})
	} else {
		messageTypes := &response.MessageTypes{
			NotFound: "liking.error.notFound"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title deleteLikingOnArticle
// @Description Delete a liking on article.
// @Accept  json
// @Param   articleId        path    int     true        "Article Id"
// @Param   id                path    int     true        "Liking Id"
// @Success 200 {object} response.BasicResponse "Liking is deleted successfully"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Liking is not deleted"
// @Resource /articles
// @Router /articles/{id}/likings/{likingId} [delete]
func deleteLikingOnArticle(c *gin.Context) {
	status, err := articleService.DeleteLikingOnArticle(c)
	messageTypes := &response.MessageTypes{
		OK:                  "liking.unlike.done",
		Unauthorized:        "liking.error.unauthorized",
		Forbidden:           "liking.error.forbidden",
		NotFound:            "liking.error.notFound",
		InternalServerError: "liking.unlike.fail"}
	messages := &response.Messages{OK: "Article liked successfully."}

	response.JSON(c, status, messageTypes, messages, err)
}
