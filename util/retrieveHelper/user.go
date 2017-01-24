package retrieveHelper

import (
	"errors"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
)

// RetriveUserWithAppIDAndUserName retrieve an user with AppID and user's name
func RetriveUserWithAppIDAndUserName(appID int64, userName string) (model.User, error) {
	var user model.User
	if db.ORM.Where("app_id=? and name=?", appID, userName).First(&user).RecordNotFound() {
		return user, errors.New("User not found. (Check the AppID and the UserName)")
	}
	return user, nil
}
