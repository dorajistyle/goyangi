package uploadService

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/aws"
	"github.com/dorajistyle/goyangi/util/file"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/stringHelper"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateFile creates a file.
func CreateFile(c *gin.Context) (int, error) {
	var form FileForm
	c.BindWith(&form, binding.Form)
	log.Debugf("CreateFile form : %v", form)
	// err = json.Unmarshal(formDataBuffer[:n], &articles)
	// file := model.File{Name: form.Name, Size: form.Size}
	user, _ := userService.CurrentUser(c)
	form.UserId = user.Id
	file := model.File{}
	modelHelper.AssignValue(&file, &form)
	if db.ORM.Create(&file).Error != nil {
		return http.StatusInternalServerError, errors.New("File is not created.")
	}
	return http.StatusCreated, nil
}

// CreateFiles creates files.
func CreateFiles(c *gin.Context) (int, error) {
	var forms FilesForm
	start := time.Now()
	c.BindWith(&forms, binding.JSON)
	log.Debugf("CreateFiles c form : %v", forms)

	user, _ := userService.CurrentUser(c)
	sqlStrBuffer := new(bytes.Buffer)
	stringHelper.Concat(sqlStrBuffer, "INSERT INTO file(user_id, name, size, created_at) VALUES ")
	values := []interface{}{}
	for _, file := range forms.Files {
		stringHelper.Concat(sqlStrBuffer, "(?, ?, ?, ?),")
		values = append(values, user.Id, file.Name, file.Size, time.Now())

	}
	// sqlStrBuffer.Truncate(sqlStrBuffer.Len() - 1) is slower than slice.
	if len(values) > 0 {
		sqlStr := sqlStrBuffer.String()
		sqlStr = sqlStr[0 : len(sqlStr)-1]
		log.Debugf("sqlStr for File : %s", sqlStr)
		db.ORM.Exec(sqlStr, values...)
		elapsed := time.Since(start)
		log.Debugf("CreateFiles elapsed : %s", elapsed)
	}

	return http.StatusCreated, nil
}

// RetrieveFile retrieves a file.
func RetrieveFile(c *gin.Context) (model.File, int, error) {
	var file model.File
	id := c.Params.ByName("id")
	if db.ORM.First(&file, id).RecordNotFound() {
		return file, http.StatusNotFound, errors.New("File is not found.")
	}
	return file, http.StatusOK, nil
}

// RetrieveFiles retrieves files.
func RetrieveFiles(c *gin.Context) []model.File {
	var files []model.File
	db.ORM.Find(&files)
	return files
}

// UpdateFile updates a file.
func UpdateFile(c *gin.Context) (model.File, int, error) {
	var file model.File
	var form FileForm
	id := c.Params.ByName("id")
	c.BindWith(&form, binding.Form)
	if db.ORM.First(&file, id).RecordNotFound() {
		return file, http.StatusNotFound, errors.New("File is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, file.UserId)
	if err != nil {
		return file, status, err
	}
	file.Name = form.Name
	file.Size = form.Size
	if db.ORM.Save(&file).Error != nil {
		return file, http.StatusInternalServerError, errors.New("File is not updated.")
	}
	return file, http.StatusOK, nil
}

// DeleteFile deletes a file.
func DeleteFile(c *gin.Context) (int, error) {
	log.Debug("deleteFile performed")
	var targetFile model.File
	id := c.Params.ByName("id")
	if db.ORM.First(&targetFile, id).RecordNotFound() {
		return http.StatusNotFound, errors.New("File is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, targetFile.UserId)
	if err != nil {
		return status, err
	}
	switch config.UploadTarget {
	case "S3":
		s3UploadPath := config.UploadS3Path + strconv.FormatInt(targetFile.UserId, 10) + "/"
		log.Debugf("s3UploadPath %s", s3UploadPath)
		err = aws.DelFromMyBucket(s3UploadPath, targetFile.Name)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	case "LOCAL":
		err = file.DeleteLocal(targetFile.Name)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	if db.ORM.Delete(&targetFile).Delete(targetFile).Error != nil {
		return http.StatusInternalServerError, errors.New("File is not deleted.")
	}
	return http.StatusOK, nil
}
