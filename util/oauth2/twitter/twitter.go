package twitter

import (
	"net/url"

	oauthHelper "github.com/dorajistyle/goyangi/util/oauth2"
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
	id, secret, redirectURL := oauthHelper.GetProvider("twitter")
	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  redirectURL,
		// Scopes: []string{
		// 	,
		// },
		Endpoint: Endpoint,
	}
}
