package route

import (
	"github.com/dorajistyle/goyangi/config"
	"github.com/gin-gonic/gin"
)

func Main(parentRoute *gin.RouterGroup) {
	route := parentRoute
	route.GET("", func(c *gin.Context) {
		sPARoute(c, "")
	})
	route.GET("/locales/:locale", func(c *gin.Context) {
		language := c.Params.ByName("locale")
		c.HTML(200, "base", map[string]string{"language": language, "staticUrl": config.StaticUrl, "guid": config.Guid, "route": "login"})
	})
}
