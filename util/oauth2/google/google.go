package google

import (
	"net/url"

	oauthHelper "github.com/dorajistyle/goyangi/util/oauth2"
	"golang.org/x/oauth2"
)

const (
	ProviderId = 1
	Scheme     = "https"
	Host       = "www.googleapis.com"
	Opaque     = "//www.googleapis.com/plus/v1/people/me"
	AuthURL    = "https://accounts.google.com/o/oauth2/auth"
	TokenURL   = "https://accounts.google.com/o/oauth2/token"
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
	id, secret, redirectURL := oauthHelper.GetProvider("google")
	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"email",
			"profile",
			"https://www.googleapis.com/auth/plus.login",
			"https://www.googleapis.com/auth/plus.profile.emails.read",
		},
		Endpoint: Endpoint,
	}
}
