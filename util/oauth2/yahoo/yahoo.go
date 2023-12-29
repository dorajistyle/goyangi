package yahoo

import (
	"net/url"

	oauthHelper "github.com/dorajistyle/goyangi/util/oauth2"
	"golang.org/x/oauth2"
)

const (
	ProviderId = 3
	Scheme     = "https"
	Host       = "social.yahooapis.com"
	Opaque     = "//social.yahooapis.com/v1/user"
	AuthURL    = "https://api.login.yahoo.com/oauth2/request_auth"
	TokenURL   = "https://api.login.yahoo.com/oauth2/get_token"
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
	id, secret, redirectURL := oauthHelper.GetProvider("yahoo")
	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"Read (Shared) Yahoo Profiles",
		},
		Endpoint: Endpoint,
	}
}
