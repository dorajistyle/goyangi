package route

import (
	"github.com/dorajistyle/goyangi/config"
	"github.com/gin-gonic/gin"
)

func sPARoute(c *gin.Context, r string) {
	switch config.Environment {
	case "DEVELOPMENT":
		c.HTML(200, "base", map[string]string{"language": config.DefaultLanguage, "title": config.Title, "staticUrl": config.StaticUrl, "guid": config.Guid, "route": r})
	default:
		c.HTML(200, "base", map[string]string{"language": config.DefaultLanguage, "title": config.Title, "staticUrl": config.StaticUrl, "guid": config.Guid, "route": r})
	}
}

func Route(route *gin.RouterGroup) {
	Admin(route)
	Articles(route)
	Locations(route)
	Pages(route)
	Setting(route)
	Upload(route)
	Users(route)
	Main(route)
}
