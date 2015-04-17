package route

import "github.com/gin-gonic/gin"

func Locations(parentRoute *gin.RouterGroup) {

	route := parentRoute.Group("/locations")
	route.GET("/", func(c *gin.Context) {
		sPARoute(c, "/locations")
	})
	route.GET("user/:username", func(c *gin.Context) {
		username := c.Params.ByName("username")
		sPARoute(c, "/locations/user/"+username)
	})
	route.GET("tag/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		sPARoute(c, "/locations/tag/"+name)
	})
	route.GET("/write", func(c *gin.Context) {
		sPARoute(c, "/locations/write")
	})
	route.GET("/write/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		sPARoute(c, "/locations/write/"+id)
	})
	route.GET("/search", func(c *gin.Context) {
		sPARoute(c, "/locations/search")
	})
	route.GET("/search/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		sPARoute(c, "/locations/search/"+id)
	})

	route = parentRoute.Group("/location")
	route.GET("/:name/:id", func(c *gin.Context) {
		name := c.Params.ByName("name")
		id := c.Params.ByName("id")
		sPARoute(c, "/location/"+name+"/"+id)
	})

}
