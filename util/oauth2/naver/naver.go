package naver

import (
	"net/url"

	oauthHelper "github.com/dorajistyle/goyangi/util/oauth2"
	"golang.org/x/oauth2"
)

const (
	ProviderId = 8
	Scheme     = "https"
	Host       = "apis.naver.com"
	Opaque     = "//apis.naver.com/nidlogin/nid/getUserProfile.xml"
	AuthURL    = "https://nid.naver.com/oauth2.0/authorize"
	TokenURL   = "https://nid.naver.com//oauth2.0/token"
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
	id, secret, redirectURL := oauthHelper.GetProvider("naver")
	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"user",
			"user:email",
		},
		Endpoint: Endpoint,
	}
}
