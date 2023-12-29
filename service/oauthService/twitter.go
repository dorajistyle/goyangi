package oauthService

import (
	// "encoding/json"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2"
	"github.com/dorajistyle/goyangi/util/oauth2/twitter"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// TwitterURL return twitter auth url.
func TwitterURL() (string, int) {
	return oauth2.OauthURL(twitter.Config), http.StatusOK
}

// SetTwitterUser set twitter user.
func SetTwitterUser(response *http.Response) (*TwitterUser, error) {
	twitterUser := &TwitterUser{}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return twitterUser, err
	}
	unmarshalErr := json.Unmarshal(body, &twitterUser)
	if unmarshalErr != nil {
		return twitterUser, unmarshalErr
	}
	log.Debugf("\ntwitterUser: %v\n", twitterUser)
	return twitterUser, err
}

// OauthTwitter link connection and user.
func OauthTwitter(c *gin.Context) (int, error) {
	var authResponse oauth2.AuthResponse
	var oauthUser OauthUser
	bindErr := c.MustBindWith(&authResponse, binding.Form)
	log.Debugf("bind error : %s\n", bindErr)
	log.Debugf("oauthRedirect form: %v", authResponse)
	response, token, err := oauth2.OauthRequest(twitter.RequestURL, twitter.Config, authResponse)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	twitterUser, err := SetTwitterUser(response)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	modelHelper.AssignValue(&oauthUser, twitterUser)
	log.Debugf("\noauthUser item : %v", oauthUser)
	status, err := LoginOrCreateOauthUser(c, &oauthUser, twitter.ProviderId, token)
	if err != nil {
		return status, err
	}
	return http.StatusSeeOther, nil
}

// RevokeTwitter revokes twitter oauth connection.
func RevokeTwitter(c *gin.Context) (map[string]bool, int, error) {
	return RevokeOauth(c, twitter.ProviderId)
}
