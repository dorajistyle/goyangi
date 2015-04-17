package github

import (
	"net/url"

	"github.com/dorajistyle/goyangi/config"
	"golang.org/x/oauth2"
)

const (
	ProviderId = 2
	Scheme     = "https"
	Host       = "api.github.com"
	Opaque     = "//api.github.com/user"
	AuthURL    = "https://github.com/login/oauth/authorize"
	TokenURL   = "https://github.com/login/oauth/access_token"
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
		ClientID:     config.OauthGithubClientID,
		ClientSecret: config.OauthGithubClientSecret,
		RedirectURL:  config.OauthGithubRedirectURL,
		Scopes: []string{
			"user",
			"user:email",
		},
		Endpoint: Endpoint,
	}
}
