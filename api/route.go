// @APIVersion 1.0.0
// @Title Goyangi API
// @Description Goyangi API usually works as expected. But sometimes its not true
// @Contact api@contact.me
// @TermsOfServiceUrl http://google.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
// @SubApi Authentication [/authentications]
// @SubApi Users [/users]
// @SubApi Oauth [/oauth]
// @SubApi Roles [/roles]
// @SubApi Articles [/articles]
// @SubApi Upload [/upload]
// @SubApi Commands [/commands]
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	v1 "github.com/dorajistyle/goyangi/api/v1"
)

// RouteAPI contains router groups for API
func RouteAPI(parentRoute *gin.Engine) {

	route := parentRoute.Group(viper.GetString("api.url"))
	{
		v1.Users(route)
		v1.Roles(route)
		v1.Authentications(route)
		v1.Articles(route)
		v1.Upload(route)
		v1.Commands(route)
		v1.Oauth(route)
	}

}
