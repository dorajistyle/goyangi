package route

import (
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/gin-gonic/gin"
)

func Users(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("")

	route.GET("/li", func(c *gin.Context) {
		userService.SetCookieHandler(c, "admin@goyangi.github.io", "password")
		sPARoute(c, "")
	})

	route.GET("/login", func(c *gin.Context) {
		sPARoute(c, "/login")
	})
	route.GET("/logout", func(c *gin.Context) {
		sPARoute(c, "/logout")
	})

	route.GET("/reset/password/:token", func(c *gin.Context) {
		token := c.Params.ByName("token")
		sPARoute(c, "/reset/password/"+token)
	})
	route.GET("/verify/email/:token", func(c *gin.Context) {
		token := c.Params.ByName("token")
		sPARoute(c, "/verify/email/"+token)
	})

	route = parentRoute.Group("/send")
	route.GET("/email/verification/form", func(c *gin.Context) {
		sPARoute(c, "/send/email/verification/form")
	})
	route.GET("/password/reset/form", func(c *gin.Context) {
		sPARoute(c, "/send/password/reset/form")
	})

	route = parentRoute.Group("/profile")
	route.GET("/", func(c *gin.Context) {
		sPARoute(c, "profile")
	})
	route.GET("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		sPARoute(c, "/profile/"+id)
	})

}
