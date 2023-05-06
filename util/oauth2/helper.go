// Helper functions of oauth2.

package oauth2

import (
	"net/url"
	"strings"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/spf13/viper"
)

func GetProvider(provider string) (id string, secret string, redirectURL string) {
	// Get provider id, secret, redirectURL from configuration.
	configRoot := []string{"oauth.", provider}
	configId := append(configRoot, ".id")
	configSecret := append(configRoot, ".secret")
	id = viper.GetString(strings.Join(configId, ""))
	secret = viper.GetString(strings.Join(configSecret, ""))
	redirectURL, err := url.JoinPath(viper.GetString("api.hostURL"), viper.GetString("api.url"), "oauth", provider, "redirect")
	if err != nil {
		log.Error("Redirect url generation failed.", err)
	}
	return id, secret, redirectURL
}
