package log

import (
	"net/http"
	"os"
	"time"

	// "github.com/dorajistyle/goyangi/util/octokit"
	"github.com/lmittmann/tint"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func LumberJackLogger(filePath string, maxSize int, maxBackups int, maxAge int) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, //days
	}
}

func InitLogToStdoutDebug() {

	// var programLevel = new(slog.LevelVar)
	// programLevel.Set()
	// h := slog.HandlerOptions{Level: programLevel}.NewTextHandler(os.Stdout) # Legacy without tint
	h := tint.Options{Level: slog.LevelDebug, TimeFormat: time.Kitchen}.NewHandler(os.Stdout)
	slog.SetDefault(slog.New(h))

}

func InitLogToStdout() {
	h := tint.Options{Level: slog.LevelWarn, TimeFormat: time.DateTime}.NewHandler(os.Stdout)
	slog.SetDefault(slog.New(h))
}

func InitLogToFile() {
	out := LumberJackLogger(viper.GetString("log.error.filepath"), viper.GetInt("log.error.maxSize"), viper.GetInt("log.error.maxBackups"), viper.GetInt("log.error.maxAge"))
	var programLevel = new(slog.LevelVar)
	programLevel.Set(slog.LevelWarn)
	h := slog.HandlerOptions{Level: programLevel}.NewJSONHandler(out)
	slog.SetDefault(slog.New(h))
}

// Init slog
func Init(environment string) {

	switch environment {
	case "DEVELOPMENT":
		InitLogToStdoutDebug()
	case "TEST":
		InitLogToFile()
	case "PRODUCTION":
		InitLogToFile()
	}
	slog.Info("", "Environment", environment)
}

// Debug logs a message with debug log level.
func Debug(msg string) {
	slog.Debug(msg)
}

// Debugf logs a formatted message with debug log level.
func Debugf(msg string, args ...interface{}) {
	slog.Debug(msg, args...)
}

// Info logs a message with info log level.
func Info(msg string) {
	slog.Info(msg)
}

// Infof logs a formatted message with info log level.
func Infof(msg string, args ...interface{}) {
	slog.Info(msg, args...)
}

// Warn logs a message with warn log level.
func Warn(msg string) {
	slog.Warn(msg)
}

// Warnf logs a formatted message with warn log level.
func Warnf(msg string, args ...interface{}) {
	slog.Warn(msg, args...)
}

// Error logs a message with error log level.
func Error(msg string, err error) {
	slog.Error(msg, err)
}

// Errorf logs a formatted message with error log level.
func Errorf(msg string, err error, args ...any) {
	slog.Error(msg, err, args)
}

// log response body data for debugging
func DebugResponse(response *http.Response) string {
	bodyBuffer := make([]byte, 5000)
	var str string
	count, err := response.Body.Read(bodyBuffer)
	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {
		if err != nil {
		}
		str += string(bodyBuffer[:count])
	}
	slog.Debug("response data : %v", str)
	return str
}
