package config

// Constants for frontend.
const (
	Title                 = "Goyangi"
	Frontend              = "CanJS"
	DefaultLanguage       = "en"
	StaticUrl             = "/static"
	StaticPathDevelopment = "frontend/canjs/static"
	StaticPathProduction  = "frontend/canjs/static-build"
)

// StaticPath return static path for each environment.
func StaticPath() string {
	var staticPath string
	switch Environment {
	case "DEVELOPMENT":
		staticPath = StaticPathDevelopment
	default:
		staticPath = StaticPathProduction
	}
	return staticPath
}
