package linkedin

import (
	"net/url"

	"github.com/dorajistyle/goyangi/config"
	"golang.org/x/oauth2"
)

const (
	ProviderId = 6
	Scheme     = "https"
	Host       = "api.linkedin.com"
	Opaque     = "//api.linkedin.com/v1/people/~:(id,first-name,email-address,picture-url,public-profile-url)?format=json"
	AuthURL    = "https://www.linkedin.com/uas/oauth2/authorization"
	TokenURL   = "https://www.linkedin.com/uas/oauth2/accessToken"
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
		ClientID:     config.OauthLinkedinClientID,
		ClientSecret: config.OauthLinkedinClientSecret,
		RedirectURL:  config.OauthLinkedinRedirectURL,
		Scopes: []string{
			"r_emailaddress",
			"r_basicprofile",
			"r_fullprofile",
		},
		Endpoint: Endpoint,
	}
}
