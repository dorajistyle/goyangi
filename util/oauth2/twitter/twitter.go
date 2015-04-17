package twitter

import (
	"net/url"

	"github.com/dorajistyle/goyangi/config"
	"golang.org/x/oauth2"
)

const (
	ProviderId = 5
	Scheme     = "https"
	Host       = "api.twitter.com"
	Opaque     = "//api.twitter.com/1.1/account/verify_credentials.json"
	AuthURL    = "https://api.twitter.com/oauth/authenticate"
	TokenURL   = "https://api.twitter.com/oauth/access_token"
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
		ClientID:     config.OauthTwitterClientID,
		ClientSecret: config.OauthTwitterClientSecret,
		RedirectURL:  config.OauthTwitterRedirectURL,
		// Scopes: []string{
		// 	,
		// },
		Endpoint: Endpoint,
	}
}
