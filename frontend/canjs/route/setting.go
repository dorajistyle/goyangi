package route

import "github.com/gin-gonic/gin"

func Setting(parentRoute *gin.RouterGroup) {
	route := parentRoute
	route = parentRoute.Group("/setting")
	route.GET("/", func(c *gin.Context) {
		sPARoute(c, "/setting")
	})
	route.GET("/:tab", func(c *gin.Context) {
		tab := c.Params.ByName("tab")
		sPARoute(c, "/setting/"+tab)
	})
}
