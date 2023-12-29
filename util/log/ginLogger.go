package log

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// GetContextLogInfo gets context log information
func GetContextLogInfo(c *gin.Context) (string, int, string, string, string) {
	method := c.Request.Method
	statusCode := c.Writer.Status()
	urlPath := c.Request.URL.Path
	errorString := c.Errors.String()
	clientIP := c.ClientIP()
	// clientIP := c.Request.RemoteAddr
	return method, statusCode, urlPath, errorString, clientIP
}

// AccessLogger write access log into a file.
func AccessLogger() gin.HandlerFunc {
	// f, err := os.OpenFile(viper.GetString("log.access.filepath"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()
	out := LumberJackLogger(viper.GetString("log.access.filepath"), viper.GetInt("log.access.maxSize"), viper.GetInt("log.access.maxBackups"), viper.GetInt("log.access.maxAge"))
	stdlogger := log.New(out, "", 0)
	// stdlogger := log.New(f, "", 0)
	// stdlogger := log.New(os.Stdout, "", 0)
	// errlogger := log.New(os.Stderr, "", 0)

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		method, statusCode, urlPath, errorString, clientIP := GetContextLogInfo(c)

		stdlogger.Printf("[GIN] %v |%3d| %12v |%s %-7s | %s | %s",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			method,
			urlPath,
			clientIP,
			errorString,
		)
	}
}

func colorForStatus(status int) string {
	switch {
	case status >= 200 && status <= 299:
		return green
	case status >= 300 && status <= 399:
		return white
	case status >= 400 && status <= 499:
		return yellow
	default:
		return red
	}
}

func colorForMethod(method string) string {
	switch {
	case method == "GET":
		return blue
	case method == "POST":
		return cyan
	case method == "PUT":
		return yellow
	case method == "DELETE":
		return red
	case method == "PATCH":
		return green
	case method == "HEAD":
		return magenta
	case method == "OPTIONS":
		return white
	default:
		return reset
	}
}

// func AccessLoggerLegacy() gin.HandlerFunc {
// 	// f, err := os.OpenFile(viper.GetString("log.access.filepath"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	// if err != nil {
// 	// 	Fatalf("error opening file: %v", err)
// 	// }
// 	// defer f.Close()
// 	out := LumberJackLogger(viper.GetString("log.access.filepath"), viper.GetInt("log.access.maxSize"), viper.GetInt("log.access.maxBackups"), viper.GetInt("log.access.maxAge"))
// 	stdlogger := log.New(out, "", 0)
// 	// stdlogger := log.New(f, "", 0)
// 	// stdlogger := log.New(os.Stdout, "", 0)
// 	// errlogger := log.New(os.Stderr, "", 0)

// 	return func(c *gin.Context) {
// 		// Start timer
// 		start := time.Now()

// 		// Process request
// 		c.Next()

// 		// Stop timer
// 		end := time.Now()
// 		latency := end.Sub(start)

// 		clientIP := c.Request.RemoteAddr
// 		method := c.Request.Method
// 		statusCode := c.Writer.Status()
// 		statusColor := colorForStatus(statusCode)
// 		methodColor := colorForMethod(method)

// 		stdlogger.Printf("[GIN] %v |%s %3d %s| %12v | %s |%s  %s %-7s %s\n%s",
// 			end.Format("2006/01/02 - 15:04:05"),
// 			statusColor, statusCode, reset,
// 			latency,
// 			clientIP,
// 			methodColor, reset, method,
// 			c.Request.URL.Path,
// 			c.Errors.String(),
// 		)
// 	}
// }
