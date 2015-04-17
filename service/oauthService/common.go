package oauthService

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2/facebook"
	"github.com/dorajistyle/goyangi/util/oauth2/github"
	"github.com/dorajistyle/goyangi/util/oauth2/google"
	"github.com/dorajistyle/goyangi/util/oauth2/kakao"
	"github.com/dorajistyle/goyangi/util/oauth2/linkedin"
	"github.com/dorajistyle/goyangi/util/oauth2/naver"
	"github.com/dorajistyle/goyangi/util/oauth2/twitter"
	"github.com/dorajistyle/goyangi/util/oauth2/yahoo"
	"github.com/dorajistyle/goyangi/util/random"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

type oauthStatusMap map[string]bool

// OauthUser is a struct for connection.
type OauthUser struct {
	Id         string
	Email      string
	Username   string
	Name       string
	ImageUrl   string
	ProfileUrl string
}

// getLoginName get login name from user. It is user's username or email.
// func getLoginName(user model.User) string {
// 	username := user.Username
// 	loginName := user.Email
// 	if len(username) > 0 {
// 		loginName = username
// 	}
// 	return loginName
// }

// RetrieveOauthStatus retrieves ouath status that provider connected or not.
func RetrieveOauthStatus(c *gin.Context) (oauthStatusMap, int, error) {
	var oauthStatus oauthStatusMap
	var connections []model.Connection
	oauthStatus = make(oauthStatusMap)
	currentUser, err := userService.CurrentUser(c)
	if err != nil {
		return oauthStatus, http.StatusUnauthorized, err
	}
	db.ORM.Model(&currentUser).Association("Connections").Find(&connections)
	for _, connection := range connections {
		log.Debugf("connection.ProviderId : %d", connection.ProviderId)
		switch connection.ProviderId {
		case google.ProviderId:
			oauthStatus["google"] = true
		case github.ProviderId:
			oauthStatus["github"] = true
		case yahoo.ProviderId:
			oauthStatus["yahoo"] = true
		case facebook.ProviderId:
			oauthStatus["facebook"] = true
		case twitter.ProviderId:
			oauthStatus["twitter"] = true
		case linkedin.ProviderId:
			oauthStatus["linkedin"] = true
		case kakao.ProviderId:
			oauthStatus["kakao"] = true
		case naver.ProviderId:
			oauthStatus["naver"] = true
		}
	}
	log.Debugf("oauthStatus : %v", oauthStatus)
	return oauthStatus, http.StatusOK, nil
}

// LoginWithOauthUser login with oauthUser's username.
func LoginWithOauthUser(c *gin.Context, token string) (int, error) {
	status, err := userService.SetCookie(c, token)
	if err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

// CreateOauthUser creates oauth user.
func CreateOauthUser(c *gin.Context, oauthUser *OauthUser, connection *model.Connection) (model.User, int, error) {
	var registrationForm userService.RegistrationForm
	var user model.User
	modelHelper.AssignValue(&registrationForm, oauthUser)
	registrationForm.Password = random.GenerateRandomString(12)
	if len(registrationForm.Username) == 0 {
		if len(registrationForm.Email) > 0 {
			registrationForm.Username = strings.Split(registrationForm.Email, "@")[0]
		} else {
			registrationForm.Username = "OauthUser"
		}
	}
	registrationForm.Username = userService.SuggestUsername(registrationForm.Username)
	generatedPassword, err := bcrypt.GenerateFromPassword([]byte(registrationForm.Password), 10)
	if err != nil {
		return user, http.StatusInternalServerError, errors.New("Password not generated.")
	}
	registrationForm.Password = string(generatedPassword)
	currentUser, err := userService.CurrentUser(c)
	if err != nil {
		log.Errorf("currentUser : %v", currentUser)
		user, err = userService.CreateUserFromForm(registrationForm)
		if err != nil {
			return user, http.StatusInternalServerError, errors.New("User is not created.")
		}
	} else {
		if db.ORM.Where("id = ?", currentUser.Id).First(&user).RecordNotFound() {
			return user, http.StatusInternalServerError, errors.New("User is not found.")
		}
	}
	db.ORM.Model(&user).Association("Connections").Append(connection)
	if db.ORM.Save(&user).Error != nil {
		return user, http.StatusInternalServerError, errors.New("Connection not appended to user.")
	}
	return user, http.StatusOK, nil
}

// LoginOrCreateOauthUser login or create with oauthUser
func LoginOrCreateOauthUser(c *gin.Context, oauthUser *OauthUser, providerID int64, token *oauth2.Token) (int, error) {
	var connection model.Connection
	var count int
	db.ORM.Where("provider_id = ? and provider_user_id = ?", providerID, oauthUser.Id).First(&connection).Count(&count)

	connection.ProviderId = providerID
	connection.ProviderUserId = oauthUser.Id
	connection.AccessToken = token.AccessToken
	connection.ProfileUrl = oauthUser.ProfileUrl
	connection.ImageUrl = oauthUser.ImageUrl
	log.Debugf("connection count : %v", count)
	if count == 1 {
		var user model.User
		if db.ORM.First(&user, "id = ?", connection.UserId).RecordNotFound() {
			return http.StatusNotFound, errors.New("User is not found.")
		}
		log.Debugf("user : %v", user)
		if db.ORM.Save(&connection).Error != nil {
			return http.StatusInternalServerError, errors.New("Connection is not updated.")
		}
		status, err := LoginWithOauthUser(c, user.Token)
		return status, err
	}
	log.Debugf("Connection is not exist.")
	user, status, err := CreateOauthUser(c, oauthUser, &connection)
	if err != nil {
		return status, err
	}
	status, err = LoginWithOauthUser(c, user.Token)
	return status, err
}

// RevokeOauth revokes oauth connection.
func RevokeOauth(c *gin.Context, providerID int64) (oauthStatusMap, int, error) {
	var oauthStatus oauthStatusMap
	var connection model.Connection
	currentUser, err := userService.CurrentUser(c)
	if err != nil {
		return oauthStatus, http.StatusUnauthorized, err
	}
	if db.ORM.First(&connection, "user_id= ? and provider_id = ? ", currentUser.Id, providerID).RecordNotFound() {
		return oauthStatus, http.StatusNotFound, errors.New("Connection is not found.")
	}
	if db.ORM.Delete(&connection).Error != nil {
		return oauthStatus, http.StatusInternalServerError, errors.New("Connection not revoked from user.")
	}
	oauthStatus, status, err := RetrieveOauthStatus(c)
	return oauthStatus, status, err
}
