// @APIVersion 1.0.0
// @Title Goyangi API
// @Description Goyangi API usually works as expected. But sometimes its not true
// @Contact api@contact.me
// @TermsOfServiceUrl http://google.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
// @SubApi Authentication [/authentications]
// @SubApi User [/users]
// @SubApi User extra [/user]
// @SubApi Oauth [/oauth]
// @SubApi Role [/roles]
// @SubApi Article [/articles]
// @SubApi Location [/locations]
// @SubApi Upload [/locations]
package api

import (
	"github.com/gin-gonic/gin"

	"github.com/dorajistyle/goyangi/api/v1"
	"github.com/dorajistyle/goyangi/config"
)

// RouteAPI contains router groups for API
func RouteAPI(parentRoute *gin.Engine) {
	route := parentRoute.Group(config.APIURL)
	{
		v1.Users(route)
		v1.Roles(route)
		v1.Authentications(route)
		v1.Articles(route)
		v1.Locations(route)
		v1.Upload(route)
		v1.Experiments(route)
		v1.Oauth(route)
	}

}
