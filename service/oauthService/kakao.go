package oauthService

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/oauth2"
	"github.com/dorajistyle/goyangi/util/oauth2/kakao"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// KakaoURL return kakao auth url.
func KakaoURL() (string, int) {
	return oauth2.OauthURL(kakao.Config), http.StatusOK
}

// SetKakaoUser set kakao user.
func SetKakaoUser(response *http.Response) (*KakaoUser, error) {
	kakaoUser := &KakaoUser{}
	defer response.Body.Close()

	log.Debugf("response.Body: %v\n", response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return kakaoUser, err
	}
	unmarshalErr := json.Unmarshal(body, &kakaoUser)
	if unmarshalErr != nil {
		return kakaoUser, unmarshalErr
	}
	log.Debugf("\nkakaoUser: %v\n", kakaoUser)
	return kakaoUser, err
}

// OauthKakao link connection and user.
func OauthKakao(c *gin.Context) (int, error) {
	var authResponse oauth2.AuthResponse
	var oauthUser OauthUser
	bindErr := c.MustBindWith(&authResponse, binding.Form)
	log.Debugf("bind error : %s\n", bindErr)
	log.Debugf("oauthRedirect form: %v", authResponse)
	response, token, err := oauth2.OauthRequest(kakao.RequestURL, kakao.Config, authResponse)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	kakaoUser, err := SetKakaoUser(response)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	modelHelper.AssignValue(&oauthUser, kakaoUser)
	log.Debugf("\noauthUser item : %v", oauthUser)
	log.Debugf("\noauthUser token : %v", token)
	status, err := LoginOrCreateOauthUser(c, &oauthUser, kakao.ProviderId, token)
	if err != nil {
		return status, err
	}
	return http.StatusSeeOther, nil
}

// RevokeKakao revokes kakao oauth connection.
func RevokeKakao(c *gin.Context) (map[string]bool, int, error) {
	return RevokeOauth(c, kakao.ProviderId)
}
