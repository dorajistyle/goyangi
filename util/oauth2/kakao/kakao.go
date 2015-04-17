package kakao

import (
	"net/url"

	"github.com/dorajistyle/goyangi/config"
	"golang.org/x/oauth2"
)

const (
	ProviderId = 7
	Scheme     = "https"
	Host       = "kapi.kakao.com"
	Opaque     = "//kapi.kakao.com/v1/user/me"
	AuthURL    = "https://kauth.kakao.com/oauth/authorize"
	TokenURL   = "https://kauth.kakao.com/oauth/token"
)

var RequestURL = &url.URL{
	Scheme: Scheme,
	Host:   Host,
	Opaque: Opaque,
}

var Endpoint = oauth2.Endpoint{
	AuthURL:  AuthURL,
	TokenURL: TokenURL,
}

var Config = Oauth2Config()

func Oauth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.OauthKakaoClientID,
		ClientSecret: config.OauthKakaoClientSecret,
		RedirectURL:  config.OauthKakaoRedirectURL,
		Scopes: []string{
			"Basic_Profile",
		},
		Endpoint: Endpoint,
	}
}
