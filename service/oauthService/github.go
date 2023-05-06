package oauthService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2"
	"github.com/dorajistyle/goyangi/util/oauth2/github"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// GithubURL return github auth url.
func GithubURL() (string, int) {
	return oauth2.OauthURL(github.Config), http.StatusOK
}

// SetGithubUser set github user.
func SetGithubUser(response *http.Response) (*GithubUser, error) {
	githubUser := &GithubUser{}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return githubUser, err
	}
	unmarshalErr := json.Unmarshal(body, &githubUser)
	if unmarshalErr != nil {
		return githubUser, unmarshalErr
	}
	log.Debugf("\ngithubUser: %v\n", githubUser)
	return githubUser, err
}

// OauthGithub link connection and user.
func OauthGithub(c *gin.Context) (int, error) {
	var authResponse oauth2.AuthResponse
	var oauthUser OauthUser
	bindErr := c.MustBindWith(&authResponse, binding.Form)
	log.Debugf("bind error : %s\n", bindErr)
	log.Debugf("oauthRedirect form: %v", authResponse)
	response, token, err := oauth2.OauthRequest(github.RequestURL, github.Config, authResponse)
	if err != nil {
		log.Error("get response error", err)
		return http.StatusInternalServerError, err
	}
	githubUser, err := SetGithubUser(response)
	if err != nil {
		log.Error("SetGithubUser error", err)
		return http.StatusInternalServerError, err
	}
	modelHelper.AssignValue(&oauthUser, githubUser)
	oauthUser.ID = strconv.Itoa(githubUser.UserID)
	log.Debugf("\noauthUser item : %v", oauthUser)
	status, err := LoginOrCreateOauthUser(c, &oauthUser, github.ProviderId, token)
	if err != nil {
		log.Errorf("LoginOrCreateOauthUser error", err, "status", status)
		return status, err
	}
	return http.StatusSeeOther, nil
}

// RevokeGithub revokes github oauth connection.
func RevokeGithub(c *gin.Context) (map[string]bool, int, error) {
	log.Debug("RevokeGithub performed")
	return RevokeOauth(c, github.ProviderId)
}
