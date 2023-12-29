package authHelper

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/gin-gonic/gin"
)

// GetAuthorizedAppFromContext gets an authorized app from *gin.Context
func GetAuthorizedAppFromContext(c *gin.Context) (model.App, int, error) {
	var app model.App
	var status int
	var err error
	_, claims, status, err := AuthenticateServer(c)
	if err != nil {
		return app, status, err
	}
	app, status, err = GetAuthorizedApp(claims["ak"], claims["sk"])
	return app, status, err
}

// CreateAuthorizedApp creates an authorized app
func CreateAuthorizedAppAndUser(appKey string, secretkey string, name string, username string) (model.App, int, error) {
	var app model.App
	var user model.User
	var err error
	result := db.ORM.First(&app, "key = ? and token = ?", appKey, secretkey)
	if result.RowsAffected == 0 {
		app.Key = appKey
		app.Token = secretkey
		app.Name = name
		db.ORM.Create(&app)
		user.Name = username
		user.AppId = app.Id
		db.ORM.Create(&user)
	}
	return app, http.StatusOK, err
}

// RemoveAuthorizedApp removes an authorized app
func RemoveAuthorizedApp(appKey string, secretkey string) (int, error) {
	result := db.ORM.Where("key = ? and token = ?", appKey, secretkey).Delete(&model.App{})
	return http.StatusOK, result.Error
}

// GetAuthorizedApp gets an authorized app from *gin.Context
func GetAuthorizedApp(appKey string, secretkey string) (model.App, int, error) {
	var app model.App
	var status int
	var err error
	if db.ORM.First(&app, "key = ? and token = ?", appKey, secretkey).RecordNotFound() {
		return app, http.StatusNotFound, errors.New("App is not found.")
	}
	return app, status, err
}
