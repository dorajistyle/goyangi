package route

import "github.com/gin-gonic/gin"

func Pages(parentRoute *gin.RouterGroup) {
	route := parentRoute
	route.GET("how-it-works", func(c *gin.Context) {
		sPARoute(c, "how-it-works")
	})
}
