package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/dorajistyle/goyangi/api/response"
	"github.com/dorajistyle/goyangi/service/uploadService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
)

// @Title Upload
// @Description Upload's router group.
func Upload(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("/upload")
	route.POST("/images", userPermission.AuthRequired(uploadImages))
	route.POST("/articles", userPermission.AuthRequired(uploadAndSyncArticles))

	route.POST("/files", userPermission.AuthRequired(createFile))
	route.POST("/files/all", userPermission.AuthRequired(createFiles))
	route.GET("/files/:id", retrieveFile)
	route.GET("/files", retrieveFiles)
	route.PUT("/files/:id", userPermission.AuthRequired(updateFile))
	route.DELETE("/files/:id", userPermission.AuthRequired(deleteFile))
}

// @Title uploadImages
// @Description upload images to storage. Request should contain multipart form data.
// @Accept  json
// @Success 201 {object} gin.H "Uploaded"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Upload failed"
// @Resource /upload/images
// @Router /upload [post]
func uploadImages(c *gin.Context) {
	status, err := uploadService.UploadImages(c)
	messageTypes := &response.MessageTypes{
		OK:                  "upload.done",
		Unauthorized:        "upload.error.unauthorized",
		InternalServerError: "upload.error.internalServerError",
	}
	messages := &response.Messages{OK: "Files uploaded successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title uploadAndSyncArticles
// @Description upload images to storage. And sync article data. Request should contain multipart form data.
// @Accept  json
// @Success 201 {object} gin.H "Uploaded"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "Upload failed"
// @Resource /upload/articles
// @Router /upload [post]
func uploadAndSyncArticles(c *gin.Context) {
	status, err := uploadService.UploadAndSyncArticles(c)
	messageTypes := &response.MessageTypes{
		OK:                  "upload.done",
		Unauthorized:        "upload.error.unauthorized",
		InternalServerError: "upload.error.internalServerError",
	}
	messages := &response.Messages{OK: "Files uploaded successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title createFile
// @Description Create a file.
// @Accept  json
// @Param   name        form   string     true        "Name of File."
// @Param   size        form   int  true        "Description of File."
// @Success 201 {object} model.File "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "File is not created"
// @Resource /upload/files
// @Router /upload [post]
func createFile(c *gin.Context) {
	status, err := uploadService.CreateFile(c)
	messageTypes := &response.MessageTypes{
		OK:                  "upload.file.create.done",
		Unauthorized:        "upload.file.error.unauthorized",
		InternalServerError: "upload.file.create.fail",
	}
	messages := &response.Messages{OK: "File model created successfully."}
	response.JSON(c, status, messageTypes, messages, err)

	// if err == nil {
	// 	c.JSON(status, gin.H{"file": file})
	// } else {
	// 	messageTypes := &response.MessageTypes{Unauthorized: "upload.file.error.unauthorized",
	// 		InternalServerError: "upload.file.create.fail"}
	// 	response.ErrorJSON(c, status, messageTypes, err)
	// }
}

// @Title createFiles
// @Description Create a file.
// @Accept  json
// @Param   name        form   string     true        "Name of File."
// @Param   size        form   int  true        "Description of File."
// @Success 201 {object} model.File "Created"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 500 {object} response.BasicResponse "File is not created"
// @Resource /upload/files
// @Router /upload [post]
func createFiles(c *gin.Context) {
	status, err := uploadService.CreateFiles(c)

	messageTypes := &response.MessageTypes{
		OK:                  "upload.file.create.done",
		Unauthorized:        "upload.file.error.unauthorized",
		InternalServerError: "upload.file.create.fail",
	}
	messages := &response.Messages{OK: "Metadata of files are created successfully."}
	response.JSON(c, status, messageTypes, messages, err)

	// if err == nil {
	// 	c.JSON(status, gin.H{"files": files})
	// } else {
	// 	messageTypes := &response.MessageTypes{Unauthorized: "upload.file.error.unauthorized",
	// 		InternalServerError: "upload.file.create.fail"}
	// 	response.ErrorJSON(c, status, messageTypes, err)
	// }
}

// @Title retrieveFile
// @Description Retrieve a file.
// @Accept  json
// @Param   id        path    int     true        "File ID"
// @Success 200 {object} model.File "OK"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /upload/files
// @Router /upload/{id} [get]
func retrieveFile(c *gin.Context) {
	file, status, err := uploadService.RetrieveFile(c)
	if err == nil {
		c.JSON(status, gin.H{"file": file})
	} else {
		messageTypes := &response.MessageTypes{NotFound: "upload.file.error.notFound"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title retrieveFiles
// @Description Retrieve file array.
// @Accept  json
// @Success 200 {array} model.File "OK"
// @Resource /upload/files
// @Router /upload [get]
func retrieveFiles(c *gin.Context) {
	files := uploadService.RetrieveFiles(c)
	c.JSON(200, gin.H{"files": files})
}

// @Title updateFile
// @Description Update a file.
// @Accept  json
// @Param   id        path    int     true        "File ID"
// @Success 200 {object} model.File "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Failure 500 {object} response.BasicResponse "File is not updated"
// @Resource /upload/files
// @Router /upload/{id} [put]
func updateFile(c *gin.Context) {
	file, status, err := uploadService.UpdateFile(c)
	if err == nil {
		c.JSON(status, gin.H{"file": file})
	} else {
		messageTypes := &response.MessageTypes{Unauthorized: "upload.file.error.unauthorized",
			NotFound:            "upload.file.error.notFound",
			InternalServerError: "upload.file.update.fail"}
		response.ErrorJSON(c, status, messageTypes, err)
	}

}

// @Title deleteFile
// @Description Delete a file.
// @Accept  json
// @Param   id        path    int     true        "File ID"
// @Success 200 {object} response.BasicResponse
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Not found"
// @Resource /upload/files
// @Router /upload/{id} [delete]
func deleteFile(c *gin.Context) {
	log.Debug("deleteFile performed")
	status, err := uploadService.DeleteFile(c)
	messageTypes := &response.MessageTypes{
		OK:           "destroy.done",
		BadRequest:   "destroy.fail",
		Unauthorized: "upload.file.error.unauthorized",
		NotFound:     "upload.file.error.notFound"}
	messages := &response.Messages{OK: "File is deleted successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}
