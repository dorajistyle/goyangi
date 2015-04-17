package oauthService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2"
	"github.com/dorajistyle/goyangi/util/oauth2/facebook"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// FacebookUser is a struct that contained facebook user information.
type FacebookUser struct {
	Id         string `json:"id"`
	Username   string `json:"name"`
	ProfileUrl string `json:"link"`
}

// FacebookURL return facebook auth url.
func FacebookURL() (string, int) {
	return oauth2.OauthURL(facebook.Config), http.StatusOK
}

// SetFacebookUser set facebook user.
func SetFacebookUser(response *http.Response) (*FacebookUser, error) {
	facebookUser := &FacebookUser{}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return facebookUser, err
	}
	json.Unmarshal(body, &facebookUser)
	log.Debugf("\nfacebookUser: %v\n", facebookUser)
	return facebookUser, nil
}

// OauthFacebook link connection and user.
func OauthFacebook(c *gin.Context) (int, error) {
	var authResponse oauth2.AuthResponse
	var oauthUser OauthUser
	c.BindWith(&authResponse, binding.Form)
	log.Debugf("oauthRedirect form: %v", authResponse)
	response, token, err := oauth2.OauthRequest(facebook.RequestURL, facebook.Config, authResponse)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	facebookUser, err := SetFacebookUser(response)
	modelHelper.AssignValue(&oauthUser, facebookUser)
	log.Debugf("\noauthUser item : %v", oauthUser)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	status, err := LoginOrCreateOauthUser(c, &oauthUser, facebook.ProviderId, token)
	if err != nil {
		return status, err
	}
	return http.StatusSeeOther, nil
}

// RevokeFacebook revokes facebook oauth connection.
func RevokeFacebook(c *gin.Context) (map[string]bool, int, error) {
	return RevokeOauth(c, facebook.ProviderId)
}
