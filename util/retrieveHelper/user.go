package retrieveHelper

import (
	"errors"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
)

// RetriveUserWithAppIdAndUserName retrieve an user with AppId and user's name
func RetriveUserWithAppIdAndUserName(appId int64, userName string) (model.User, error) {
	var user model.User
	if db.ORM.Where("app_id=? and name=?", appId, userName).First(&user).RecordNotFound() {
		return user, errors.New("User not found. (Check the AppId and the UserName)")
	}
	return user, nil
}
