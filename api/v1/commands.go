package v1

import (
	"github.com/dorajistyle/goyangi/db"
	"github.com/gin-gonic/gin"
)

// @Title Commands
// @Description Commands' router group. It contains command API.
func Commands(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("/commands")
	route.GET("migrate", migrate)
}

func migrate(c *gin.Context) {
	db.Migrate()
}