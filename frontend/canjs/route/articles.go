package route

import "github.com/gin-gonic/gin"

func Articles(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("")
	route.GET("/notice", func(c *gin.Context) {
		sPARoute(c, "/notice")
	})
	route.GET("/general", func(c *gin.Context) {
		sPARoute(c, "/general")
	})
	route.GET("/etc", func(c *gin.Context) {
		sPARoute(c, "/etc")
	})
	route.GET("/noticeOne", func(c *gin.Context) {
		sPARoute(c, "/noticeOne")
	})

	route = parentRoute.Group("/articles")
	route.GET("/", func(c *gin.Context) {
		sPARoute(c, "/articles")
	})
	route.GET("user/:username", func(c *gin.Context) {
		username := c.Params.ByName("username")
		sPARoute(c, "/articles/user/"+username)
	})
	route.GET("tag/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		sPARoute(c, "/articles/tag/"+name)
	})
	route.GET("/write", func(c *gin.Context) {
		sPARoute(c, "/articles/write")
	})
	route.GET("/notice/write", func(c *gin.Context) {
		sPARoute(c, "/articles/notice/write")
	})
	route.GET("/general/write", func(c *gin.Context) {
		sPARoute(c, "/articles/general/write")
	})
	route.GET("/etc/write", func(c *gin.Context) {
		sPARoute(c, "/articles/etc/write")
	})
	route.GET("/write/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		sPARoute(c, "/articles/write/"+id)
	})

	route = parentRoute.Group("/article")
	route.GET("/:title/:id", func(c *gin.Context) {
		title := c.Params.ByName("title")
		id := c.Params.ByName("id")
		sPARoute(c, "/article/"+title+"/"+id)
	})

}
