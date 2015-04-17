package facebook

import (
	"net/url"

	"github.com/dorajistyle/goyangi/config"
	"golang.org/x/oauth2"
)

const (
	ProviderId = 4
	Scheme     = "https"
	Host       = "graph.facebook.com"
	Opaque     = "//graph.facebook.com/me"
	AuthURL    = "https://www.facebook.com/dialog/oauth"
	TokenURL   = "https://graph.facebook.com/oauth/access_token"
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
		ClientID:     config.OauthFacebookClientID,
		ClientSecret: config.OauthFacebookClientSecret,
		RedirectURL:  config.OauthFacebookRedirectURL,
		Scopes: []string{
			"public_profile",
			"email",
		},
		Endpoint: Endpoint,
	}
}
