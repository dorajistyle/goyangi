package authHelper

import (
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/userHelper"
	"github.com/gin-gonic/gin"
)

// GetAuthorizedUserFromContext gets an authorized user from *gin.Context
func GetAuthorizedUserFromContext(c *gin.Context) (model.User, int, error) {
	var user model.User
	var status int
	var err error
	_, claims, status, err := AuthenticateServer(c)
	if err != nil {
		return user, status, err
	}
	user, status, err = GetAuthorizedUser(claims["ak"], claims["sk"], claims["un"])
	return user, status, err
}

// GetAuthorizedUser gets an authorized user from *gin.Context
func GetAuthorizedUser(appKey string, secretkey string, userName string) (model.User, int, error) {
	var app model.App
	var user model.User
	var status int
	var err error
	app, status, err = GetAuthorizedApp(appKey, secretkey)
	if err != nil {
		return user, status, err
	}
	appId := app.Id
	user, status, err = userHelper.FindUserByUserName(appId, userName)
	return user, status, err
}
