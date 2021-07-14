package log

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/dorajistyle/goyangi/util/config"
	// "github.com/dorajistyle/goyangi/util/octokit"
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
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func InitLogToStdout() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}

func InitLogToFile() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// f, err := os.OpenFile(config.ErrorLogFilePath+config.ErrorLogFileExtension, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	Fatalf("error opening file: %v", err)
	// }
	out := LumberJackLogger(config.ErrorLogFilePath+config.ErrorLogFileExtension, config.ErrorLogMaxSize, config.ErrorLogMaxBackups, config.ErrorLogMaxAge)
	// defer f.Close()
	// logrus.SetOutput(os.Stderr)
	logrus.SetOutput(out)
	logrus.SetLevel(logrus.WarnLevel)
}

// DEPRECATED BEACAUSE OF code.google.com/p/go-netrc/netrc
// func InitLogToGithub() {
// 	logrus.SetFormatter(&logrus.TextFormatter{})
// 	logrus.AddHook(octokit.NewOctokitHook(config.OctokitGitHubAPIURL,
// 		config.OctokitUserAgent,
// 		config.OctokitAccessToken,
// 		config.OctokitTargetOwner,
// 		config.OctokitTargetRepo))
//     logrus.SetOutput(os.Stdout)
// 	logrus.SetLevel(logrus.WarnLevel)
// }

// func InitLogToGithubAndFile() {
// 	logrus.SetFormatter(&logrus.TextFormatter{})
// 	logrus.AddHook(octokit.NewOctokitHook(config.OctokitGitHubAPIURL,
// 		config.OctokitUserAgent,
// 		config.OctokitAccessToken,
// 		config.OctokitTargetOwner,
// 		config.OctokitTargetRepo))
// 	out := LumberJackLogger(config.ErrorLogFilePath+config.ErrorLogFileExtension, config.ErrorLogMaxSize, config.ErrorLogMaxBackups, config.ErrorLogMaxAge)
// 	logrus.SetOutput(out)
// 	logrus.SetLevel(logrus.WarnLevel)
// }

// Init logrus
func Init(environment string) {
	// Use the Airbrake hook to report errors that have Error severity or above to
	// an exception tracker. You can create custom hooks, see the Hooks section.
	//  logrus.AddHook(&log_airbrake.AirbrakeHook{})
	switch environment {
	case "DEVELOPMENT":
		InitLogToStdoutDebug()
	case "TEST":
		InitLogToFile()
	case "PRODUCTION":
		InitLogToFile()
	}
	logrus.Debugf("Environment : %s", environment)
}

// Debug logs a message with debug log level.
func Debug(msg string) {
	logrus.Debug(msg)
}

// Debugf logs a formatted message with debug log level.
func Debugf(msg string, args ...interface{}) {
	logrus.Debugf(msg, args...)
}

// Info logs a message with info log level.
func Info(msg string) {
	logrus.Info(msg)
}

// Infof logs a formatted message with info log level.
func Infof(msg string, args ...interface{}) {
	logrus.Infof(msg, args...)
}

// Warn logs a message with warn log level.
func Warn(msg string) {
	logrus.Warn(msg)
}

// Warnf logs a formatted message with warn log level.
func Warnf(msg string, args ...interface{}) {
	logrus.Warnf(msg, args...)
}

// Error logs a message with error log level.
func Error(msg string) {
	logrus.Error(msg)
}

// Errorf logs a formatted message with error log level.
func Errorf(msg string, args ...interface{}) {
	logrus.Errorf(msg, args...)
}

// Fatal logs a message with fatal log level.
func Fatal(msg string) {
	logrus.Fatal(msg)
}

// Fatalf logs a formatted message with fatal log level.
func Fatalf(msg string, args ...interface{}) {
	logrus.Fatalf(msg, args...)
}

// Panic logs a message with panic log level.
func Panic(msg string) {
	logrus.Panic(msg)
}

// Panicf logs a formatted message with panic log level.
func Panicf(msg string, args ...interface{}) {
	logrus.Panicf(msg, args...)
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
	Debugf("response data : %v", str)
	return str
}
