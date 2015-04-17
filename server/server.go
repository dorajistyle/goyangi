package server

import (
	"github.com/dorajistyle/goyangi/api"
	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/frontend/canjs"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
)

func initI18N() {
	i18n.MustLoadTranslationFile("service/userService/locale/en-us.all.json")
	i18n.MustLoadTranslationFile("service/userService/locale/ko-kr.all.json")
}

func init() {
	log.Init()
	initI18N()
}

// CORSMiddleware for CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			// c.Abort(200)
			c.Abort()
			return
		}
		c.Next()
	}
}

func Run() {
	r := gin.New()

	// Global middlewares
	// If use gin.Logger middlewares, it send duplicated request.
	switch config.Environment {
	case "DEVELOPMENT":
		r.Use(gin.Logger())
	case "TEST":
		r.Use(log.AccessLogger())
	case "PRODUCTION":
		r.Use(log.AccessLogger())
	}
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	switch config.Frontend {
	case "CanJS":
		canjs.LoadPage(r)
	default:
		canjs.LoadPage(r)
	}
	api.RouteAPI(r)

	// Listen and server on 0.0.0.0:3001
	//    r.Run("localhost:3001")
	r.Run(":3001")
}
