package oauthService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2"
	"github.com/dorajistyle/goyangi/util/oauth2/google"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Email is a struct that contained googleUser's email.
type Email struct {
	Value string `json:"value" binding:"required"`
	Type  string `json:"type" binding:"required"`
}

// Image is a struct that contained googleUser's image meta data.
type Image struct {
	URL       string `json:"url" binding:"required"`
	IsDefault string `json:"isDefault" binding:"required"`
}

// GoogleUser is a struct that contained google user information.
type GoogleUser struct {
	Id         string  `json:"id"`
	Username   string  `json:"nickname"`
	Emails     []Email `json:"emails"`
	Name       string  `json:"displayName"`
	Image      Image   `json:"image"`
	ProfileUrl string  `json:"url"`
}

// GoogleURL return google auth url.
func GoogleURL() (string, int) {
	return oauth2.OauthURL(google.Config), http.StatusOK
}

// SetGoogleUser set google user.
func SetGoogleUser(response *http.Response) (*GoogleUser, error) {
	googleUser := &GoogleUser{}
	log.Debugf("\nresponse: %v\n", response)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return googleUser, err
	}
	json.Unmarshal(body, &googleUser)
	return googleUser, err
}

// OauthGoogle link connection and user.
func OauthGoogle(c *gin.Context) (int, error) {
	log.Debugf("c.Request.URL : %s", c.Request.URL)
	var authResponse oauth2.AuthResponse
	var oauthUser OauthUser
	c.BindWith(&authResponse, binding.Form)
	log.Debugf("oauthRedirect form: %v", authResponse)
	response, token, err := oauth2.OauthRequest(google.RequestURL, google.Config, authResponse)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	googleUser, err := SetGoogleUser(response)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	modelHelper.AssignValue(&oauthUser, googleUser)
	oauthUser.Email = googleUser.Emails[0].Value
	oauthUser.ImageUrl = googleUser.Image.URL
	status, err := LoginOrCreateOauthUser(c, &oauthUser, google.ProviderId, token)
	if err != nil {
		return status, err
	}
	return http.StatusSeeOther, nil
}

// RevokeGoogle revokes google oauth connection.
func RevokeGoogle(c *gin.Context) (map[string]bool, int, error) {
	return RevokeOauth(c, google.ProviderId)
}
