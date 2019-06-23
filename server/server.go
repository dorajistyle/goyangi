package server

import (
	"time"

	"github.com/dorajistyle/goyangi/api"
	"github.com/dorajistyle/goyangi/config"
	// "github.com/dorajistyle/goyangi/frontend/vuejs"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
)

func initI18N() {
	i18n.MustLoadTranslationFile("service/userService/locale/en-us.all.json")
	i18n.MustLoadTranslationFile("service/userService/locale/ko-kr.all.json")
}

func init() {
	log.Init(config.Environment)
	initI18N()
}

// CORSMiddleware for CORS
func CORSMiddleware() gin.HandlerFunc {
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	})
	return corsConfig
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

	api.RouteAPI(r)

	// Listen and server on 0.0.0.0:3001
	//    r.Run("localhost:3001")
	r.Run(":3001")
}
