package oauthService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2"
	"github.com/dorajistyle/goyangi/util/oauth2/naver"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// NaverUser is a struct that contained naver user information.
type NaverUser struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"login"`
	Name       string `json:"name""`
	ImageUrl   string `json:"avatar_url"`
	ProfileUrl string `json:"html_url"`
}

// NaverURL return naver auth url.
func NaverURL() (string, int) {
	return oauth2.OauthURL(naver.Config), http.StatusOK
}

// SetNaverUser set naver user.
func SetNaverUser(response *http.Response) (*NaverUser, error) {
	naverUser := &NaverUser{}
	defer response.Body.Close()
	log.Debugf("response.Body: %v\n", response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return naverUser, err
	}
	json.Unmarshal(body, &naverUser)
	log.Debugf("\nnaverUser: %v\n", naverUser)
	return naverUser, err
}

// OauthNaver link connection and user.
func OauthNaver(c *gin.Context) (int, error) {
	var authResponse oauth2.AuthResponse
	var oauthUser OauthUser
	c.BindWith(&authResponse, binding.Form)
	log.Debugf("oauthRedirect form: %v", authResponse)
	response, token, err := oauth2.OauthRequest(naver.RequestURL, naver.Config, authResponse)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	naverUser, err := SetNaverUser(response)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	modelHelper.AssignValue(&oauthUser, naverUser)
	log.Debugf("\noauthUser item : %v", oauthUser)
	log.Debugf("\noauthUser token : %v", token)
	status, err := LoginOrCreateOauthUser(c, &oauthUser, naver.ProviderId, token)
	if err != nil {
		return status, err
	}
	return http.StatusSeeOther, nil
}

// RevokeNaver revokes naver oauth connection.
func RevokeNaver(c *gin.Context) (map[string]bool, int, error) {
	return RevokeOauth(c, naver.ProviderId)
}
