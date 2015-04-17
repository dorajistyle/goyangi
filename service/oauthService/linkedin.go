package oauthService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2"
	"github.com/dorajistyle/goyangi/util/oauth2/linkedin"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// LinkedinUser is a struct that contained linkedin user information.
type LinkedinUser struct {
	Id         string `json:"id"`
	Email      string `json:"emailAddress"`
	Username   string `json:"firstName"`
	Name       string `json:"lastName""`
	ImageUrl   string `json:"pictureUrl"`
	ProfileUrl string `json:"publicProfileUrl"`
}

// LinkedinURL return linkedin auth url.
func LinkedinURL() (string, int) {
	return oauth2.OauthURL(linkedin.Config), http.StatusOK
}

// SetLinkedinUser set linkedin user.
func SetLinkedinUser(response *http.Response) (*LinkedinUser, error) {
	linkedinUser := &LinkedinUser{}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return linkedinUser, err
	}
	json.Unmarshal(body, &linkedinUser)
	log.Debugf("\nlinkedinUser: %v\n", linkedinUser)
	return linkedinUser, err
}

// OauthLinkedin link connection and user.
func OauthLinkedin(c *gin.Context) (int, error) {
	var authResponse oauth2.AuthResponse
	var oauthUser OauthUser
	c.BindWith(&authResponse, binding.Form)
	log.Debugf("oauthRedirect form: %v", authResponse)
	response, token, err := oauth2.OauthRequest(linkedin.RequestURL, linkedin.Config, authResponse)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	linkedinUser, err := SetLinkedinUser(response)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	modelHelper.AssignValue(&oauthUser, linkedinUser)
	log.Debugf("\noauthUser item : %v", oauthUser)
	log.Debugf("\nlinkedinUser id : %s", linkedinUser.Id)
	log.Debugf("\noauthUser id : %s", oauthUser.Id)
	status, err := LoginOrCreateOauthUser(c, &oauthUser, linkedin.ProviderId, token)
	if err != nil {
		return status, err
	}
	return http.StatusSeeOther, nil
}

// RevokeLinkedin revokes linkedin oauth connection.
func RevokeLinkedin(c *gin.Context) (map[string]bool, int, error) {
	return RevokeOauth(c, linkedin.ProviderId)
}
