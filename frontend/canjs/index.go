package canjs

import (
	"github.com/gin-gonic/gin"

	"html/template"
	//"net/http"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/frontend/canjs/route"
	"github.com/dorajistyle/goyangi/util/log"
)

var staticPath string

func init() {
	staticPath = config.StaticPath()
}

func LoadPage(parentRoute *gin.Engine) {
	//type Page struct {
	//    Title string
	//}

	parentRoute.SetHTMLTemplate(template.Must(template.ParseFiles("frontend/canjs/templates/message.html", "frontend/canjs/templates/app.html", "frontend/canjs/templates/base.html", "frontend/canjs/templates/404.html")))
	log.Debug("url : " + config.StaticUrl)
	log.Debug("guid : " + config.Guid)
	log.Debug("path : " + staticPath)
	parentRoute.Static(config.StaticUrl+"/"+config.Guid, staticPath)
	// route.ServeFiles doesn't exist in the current version of gin. If you want to use this, use the 59d949d35080b83864dbeafadecef112d46aaeee.
	//parentRoute.ServeFiles(config.StaticUrl+"/"+config.Guid+"/*filepath", http.Dir(staticPath))
	parentRoute.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", map[string]string{"language": config.DefaultLanguage, "title": config.Title})
	})

	route.Route(parentRoute.Group(""))
}
