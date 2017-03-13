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

// GetAuthorizedApp gets an authorized app from *gin.Context
func GetAuthorizedApp(appKey string, secretkey string) (model.App, int, error) {
	var app model.App
	var status int
	var err error
	if db.ORM.First(&app, "key = ? and secret_key = ?", appKey, secretkey).RecordNotFound() {
		return app, http.StatusNotFound, errors.New("App is not found.")
	}
	return app, status, err
}
