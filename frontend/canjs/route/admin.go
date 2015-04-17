package route

import (
	"github.com/dorajistyle/goyangi/config"
	"github.com/gin-gonic/gin"
)

func Admin(parentRoute *gin.RouterGroup) {
	route := parentRoute
	route.GET("/log/access", func(c *gin.Context) {
		c.File(config.AccessLogFilePath + config.AccessLogFileExtension)
	})
	route.GET("log/error", func(c *gin.Context) {
		c.File(config.ErrorLogFilePath + config.ErrorLogFileExtension)
	})

	route = parentRoute.Group("/admin")
	route.GET("/:type", func(c *gin.Context) {
		typeName := c.Params.ByName("type")
		sPARoute(c, "/admin/"+typeName)
	})
	route.GET("/:type/:page", func(c *gin.Context) {
		typeName := c.Params.ByName("type")
		page := c.Params.ByName("page")
		sPARoute(c, "/admin/"+typeName+"/"+page)
	})
}
