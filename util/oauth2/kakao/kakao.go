package kakao

import (
	"net/url"

	oauthHelper "github.com/dorajistyle/goyangi/util/oauth2"
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
	id, secret, redirectURL := oauthHelper.GetProvider("kakao")
	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"Basic_Profile",
		},
		Endpoint: Endpoint,
	}
}
