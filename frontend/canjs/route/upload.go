package route

import "github.com/gin-gonic/gin"

func Upload(parentRoute *gin.RouterGroup) {
	route := parentRoute
	route.GET("upload", func(c *gin.Context) {
		sPARoute(c, "upload")
	})
}
