package userHelper

import (
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"

	//   "github.com/gin-gonic/gin"
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/retrieveHelper"
	//   "github.com/dorajistyle/goyangi/util/crypto"
)

// FindUserByUserName creates a user.
func FindUserByUserName(appId int64, userName string) (model.User, int, error) {
	var user model.User
	var err error
	// token := c.Request.Header.Get("X-Auth-Token")
	user, err = retrieveHelper.RetriveUserWithAppIdAndUserName(appId, userName)
	if err != nil {
		return user, http.StatusUnauthorized, err
	}
	return user, http.StatusOK, nil
}

// FindOrCreateUser creates a user.
func FindOrCreateUser(appId int64, userName string) (model.User, int, error) {
	var user model.User
	var err error

	// if len(token) > 0 {
	// 	log.Debug("header token exist.")
	// } else {
	// 	token, err = Token(c)
	// 	log.Debug("header token not exist.")
	// 	if err != nil {
	// 		return user, http.StatusUnauthorized, err
	// 	}
	// }
	log.Debugf("userName : %s\n", userName)
	// log.Debugf("Error : %s\n", err.Error())
	user, err = retrieveHelper.RetriveUserWithAppIdAndUserName(appId, userName)
	if err != nil {
		var user model.User
		// return user, http.StatusBadRequest, err
		user.Name = userName
		// user.Token = token
		user.AppId = uint(appId)
		log.Debugf("user %+v\n", user)
		if db.ORM.Create(&user).Error != nil {
			return user, http.StatusBadRequest, errors.New("User is not created.")
		}
		log.Debugf("retrived User %v\n", user)
		return user, http.StatusOK, nil
	}
	return user, http.StatusBadRequest, nil
}
