package v1

import (
	"github.com/dorajistyle/goyangi/api/response"
	"github.com/dorajistyle/goyangi/service/locationService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/gin-gonic/gin"
)

// @Title Locations
// @Description Locations's router group.
func Locations(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("/locations")
	route.POST("/", userPermission.AuthRequired(createLocation))
	route.GET("/:id", retrieveLocation)
	route.GET("/", retrieveLocations)
	route.PUT("/:id", userPermission.AuthRequired(updateLocation))
	route.DELETE("/:id", userPermission.AuthRequired(deleteLocation))

	route.POST("/comments", userPermission.AuthRequired(createCommentOnLocation))
	route.GET("/:id/comments", retrieveCommentsOnLocation)
	route.PUT("/:id/comments/:commentId", userPermission.AuthRequired(updateCommentOnLocation))
	route.DELETE("/:id/comments/:commentId", userPermission.AuthRequired(deleteCommentOnLocation))

	route.POST("/likings", userPermission.AuthRequired(createLikingOnLocation))
	route.GET("/:id/likings", retrieveLikingsOnLocations)
	route.DELETE("/:id/likings/:userId", userPermission.AuthRequired(deleteLikingOnLocation))
}

// @Title createLocation
// @Description Create an location.
// @Accept  json
// @Param   title        form   string     true        "Location title."
// @Param   url        form   string     true        "Location url"
// @Param   latitude        form   int     true        "Location latitude"
// @Param   longitude        form   int     true        "Location longitude"
// @Param   content        form   string  false        "Location content"
// @Success 201 {object} model.Location "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Location is not created"
// @Resource /locations
// @Router /locations [post]
func createLocation(c *gin.Context) {
	location, status, err := locationService.CreateLocation(c)
	if err == nil {
		c.JSON(status, gin.H{"location": location})
	} else {
		messageTypes := &response.MessageTypes{InternalServerError: "location.error.notCreated"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title retrieveLocation
// @Description Retrieve an location.
// @Accept  json
// @Param   id        path    int     true        "Location Id"
// @Success 200 {object} model.Location "OK"
// @Failure 404 {object} response.BasicResponse "Location is not found"
// @Resource /locations
// @Router /locations/{id} [get]
func retrieveLocation(c *gin.Context) {
	location, isAuthor, currentUserId, status, err := locationService.RetrieveLocation(c)
	if err == nil {
		c.JSON(status, gin.H{"location": location, "isAuthor": isAuthor, "currentUserId": currentUserId})
	} else {
		messageTypes := &response.MessageTypes{
			NotFound: "location.error.notFound",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title retrieveLocations
// @Description Retrieve location array.
// @Accept  json
// @Success 200 {array} model.Location "OK"
// @Resource /locations
// @Router /locations [get]
func retrieveLocations(c *gin.Context) {
	locations, canWrite, currentPage, hasPrev, hasNext, status, err := locationService.RetrieveLocations(c)
	if err == nil {
		c.JSON(status, gin.H{"locations": locations, "canWrite": canWrite, "currentPage": currentPage,
			"hasPrev": hasPrev, "hasNext": hasNext})
	} else {
		messageTypes := &response.MessageTypes{}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title updateLocation
// @Description Update an location.
// @Accept  json
// @Param   id        path    int     true        "Location Id"
// @Success 200 {object} model.Location "OK"
// @Failure 400 {object} response.BasicResponse "Location is not updated"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Location is not found"
// @Resource /locations
// @Router /locations/{id} [put]
func updateLocation(c *gin.Context) {
	location, status, err := locationService.UpdateLocation(c)
	if err == nil {
		c.JSON(status, gin.H{"location": location})
	} else {
		messageTypes := &response.MessageTypes{
			BadRequest:   "location.view.updated.fail",
			Unauthorized: "location.error.isNotAuthor",
			NotFound:     "location.error.notFound"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title deleteLocation
// @Description Delete an location.
// @Accept  json
// @Param   id        path    int     true        "Location Id"
// @Success 200 {object} response.BasicResponse "Location deleted"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /locations
// @Router /locations/{id} [delete]
func deleteLocation(c *gin.Context) {
	status, err := locationService.DeleteLocation(c)
	messageTypes := &response.MessageTypes{
		OK:           "location.view.deleted.done",
		BadRequest:   "location.view.deleted.fail",
		Unauthorized: "location.error.isNotAuthor",
		NotFound:     "location.error.notFound"}
	messages := &response.Messages{OK: "Location is deleted successfully."}
	response.JSON(c, status, messageTypes, messages, err)

}

// @Title createCommentOnLocation
// @Description Create a comment on an location.
// @Accept  json
// @Param   locationId        form   int     true        "Location id."
// @Param   content        form   string     true        "Comment content."
// @Param   description        form   string  false        "Location description."
// @Success 201 {object} response.BasicResponse "Comment created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Location is not found"
// @Resource /locations
// @Router /locations/comments [post]
func createCommentOnLocation(c *gin.Context) {
	status, err := locationService.CreateCommentOnLocation(c)
	messageTypes := &response.MessageTypes{
		Created:      "comment.created.done",
		Unauthorized: "comment.error.unauthorized",
		Forbidden:    "comment.error.forbidden",
		NotFound:     "comment.error.notFound"}
	messages := &response.Messages{OK: "Comment is created successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveCommentsOnLocation
// @Description Retrieve comments on an location.
// @Accept  json
// @Param   locationId        path    int     true        "Location Id"
// @Success 200 {object} model.Comment "Comment updated successfully"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Comment is not updated"
// @Resource /locations
// @Router /locations/{id}/comments [get]
func retrieveCommentsOnLocation(c *gin.Context) {
	comments, currentPage, hasPrev, hasNext, count, status, err := locationService.RetrieveCommentsOnLocation(c)
	if err == nil {
		c.JSON(status, gin.H{"comments": comments, "currentPage": currentPage, "hasPrev": hasPrev, "hasNext": hasNext, "count": count})
	} else {
		messageTypes := &response.MessageTypes{
			NotFound: "comment.error.notFound"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title updateCommentOnLocation
// @Description Update a comment on location.
// @Accept  json
// @Param   locationId        path    int     true        "Location Id"
// @Param   id                path    int     true        "Comment Id"
// @Success 200 {object} model.Comment "Comment updated successfully"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Comment is not updated"
// @Resource /locations
// @Router /locations/{id}/comments/{commentId} [put]
func updateCommentOnLocation(c *gin.Context) {
	status, err := locationService.UpdateCommentOnLocation(c)
	messageTypes := &response.MessageTypes{
		OK:                  "comment.updated.done",
		Unauthorized:        "comment.error.unauthorized",
		Forbidden:           "comment.error.forbidden",
		NotFound:            "comment.error.notFound",
		InternalServerError: "comment.updated.fail"}
	messages := &response.Messages{OK: "Comment is created successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title deleteCommentOnLocation
// @Description Delete a comment on location.
// @Accept  json
// @Param   locationId        path    int     true        "Location Id"
// @Param   id                path    int     true        "Comment Id"
// @Success 200 {object} response.BasicResponse "Comment is deleted successfully"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Comment is not deleted"
// @Resource /locations
// @Router /locations/{id}/comments/{commentId} [delete]
func deleteCommentOnLocation(c *gin.Context) {
	status, err := locationService.DeleteCommentOnLocation(c)
	messageTypes := &response.MessageTypes{
		OK:                  "comment.deleted.done",
		Unauthorized:        "comment.error.unauthorized",
		Forbidden:           "comment.error.forbidden",
		NotFound:            "comment.error.notFound",
		InternalServerError: "comment.deleted.fail"}
	messages := &response.Messages{OK: "Comment is deleted successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title createLikingOnLocation
// @Description Create a liking on an location.
// @Accept  json
// @Param   locationId        form   int     true        "Location id."
// @Param   content        form   string     true        "Liking content."
// @Param   imageName        form   string     true        "Location image name."
// @Param   description        form   string  false        "Location description."
// @Success 201 {object} response.BasicResponse "Liking created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Location is not found"
// @Resource /locations
// @Router /locations/likings [post]
func createLikingOnLocation(c *gin.Context) {
	status, err := locationService.CreateLikingOnLocation(c)
	messageTypes := &response.MessageTypes{
		OK:           "liking.like.done",
		BadRequest:   "liking.like.fail",
		Unauthorized: "liking.error.unauthorized",
		NotFound:     "liking.error.notFound"}
	messages := &response.Messages{OK: "Location liking is created successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title retrieveLikingsOnLocations
// @Description Retrieve likings on an location.
// @Accept  json
// @Param   locationId        path    int     true        "Location Id"
// @Success 200 {array} model.PublicUser "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /locations
// @Router /locations/{id}/likings [get]
func retrieveLikingsOnLocations(c *gin.Context) {
	likings, currentPage, hasPrev, hasNext, count, status, err := locationService.RetrieveLikingsOnLocations(c)
	if err == nil {
		c.JSON(status, gin.H{"likings": likings, "currentPage": currentPage,
			"hasPrev": hasPrev, "hasNext": hasNext, "count": count})
	} else {
		messageTypes := &response.MessageTypes{
			NotFound: "liking.error.notFound"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title deleteLikingOnLocation
// @Description Delete a liking on location.
// @Accept  json
// @Param   locationId        path    int     true        "Location Id"
// @Param   id                path    int     true        "Liking Id"
// @Success 200 {object} response.BasicResponse "Liking is deleted successfully"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 403 {object} response.BasicResponse "FormUser's Id is not identical with currentUser's Id"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "Liking is not deleted"
// @Resource /locations
// @Router /locations/{id}/likings/{likingId} [delete]
func deleteLikingOnLocation(c *gin.Context) {
	status, err := locationService.DeleteLikingOnLocation(c)
	messageTypes := &response.MessageTypes{
		OK:                  "liking.unlike.done",
		Unauthorized:        "liking.error.unauthorized",
		Forbidden:           "liking.error.forbidden",
		NotFound:            "liking.error.notFound",
		InternalServerError: "liking.unlike.fail"}
	messages := &response.Messages{OK: "Article liked successfully."}

	response.JSON(c, status, messageTypes, messages, err)
}
