package oauthService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2"
	"github.com/dorajistyle/goyangi/util/oauth2/yahoo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// YahooUser is a struct that contained yahoo user information.
type YahooUser struct {
	Id         string `json:"guid"`
	Email      string `json:"email"`
	Name       string `json:"nickname"`
	Username   string `json:"nickname"`
	ImageUrl   string `json:"imageURL"`
	ProfileUrl string `json:"profileURL"`
}

// YahooURL return yahoo auth url.
func YahooURL() (string, int) {
	return oauth2.OauthURL(yahoo.Config), http.StatusOK
}

// SetYahooUser set yahoo user.
func SetYahooUser(response *http.Response) (*YahooUser, error) {
	yahooUser := &YahooUser{}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return yahooUser, err
	}
	json.Unmarshal(body, &yahooUser)
	log.Debugf("\nyahooUser: %v\n", yahooUser)
	return yahooUser, err
}

// OauthYahoo link connection and user.
func OauthYahoo(c *gin.Context) (int, error) {
	var authResponse oauth2.AuthResponse
	var oauthUser OauthUser
	c.BindWith(&authResponse, binding.Form)
	log.Debugf("oauthRedirect form: %v", authResponse)
	response, token, err := oauth2.OauthRequest(yahoo.RequestURL, yahoo.Config, authResponse)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	yahooUser, err := SetYahooUser(response)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	modelHelper.AssignValue(&oauthUser, yahooUser)
	log.Debugf("\noauthUser item : %v", oauthUser)
	status, err := LoginOrCreateOauthUser(c, &oauthUser, yahoo.ProviderId, token)
	if err != nil {
		return status, err
	}
	return http.StatusSeeOther, nil
}

// RevokeYahoo revokes yahoo oauth connection.
func RevokeYahoo(c *gin.Context) (map[string]bool, int, error) {
	return RevokeOauth(c, yahoo.ProviderId)
}
