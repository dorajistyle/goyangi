// DEPRECATED BEACAUSE OF code.google.com/p/go-netrc/netrc 
package octokit

// import (
// 	"github.com/Sirupsen/logrus"
// )

// // OctokitHook to send logs via octokit.
// type OctokitHook struct {}

// func NewOctokitHook(gitHubAPIURL, userAgent, accessToken, targeOwner, targetRepo string) (*OctokitHook) {
// 	NewOctokitClient(gitHubAPIURL, userAgent, accessToken, targeOwner, targetRepo)
// 	return &OctokitHook{}
// }

// func (hook *OctokitHook) Fire(entry *logrus.Entry) error {
// 	body := entry.Message
// 	title := body
// 	if len(body) > 60 {
// 				title = body[:60]
// 		}

// 	switch entry.Level {
// 	case logrus.PanicLevel:
// 		return SendLog(title, body, "panic")
// 	case logrus.FatalLevel:
// 		return SendLog(title, body, "fatal")
// 	case logrus.ErrorLevel:
// 		return SendLog(title, body, "error")
// 	case logrus.WarnLevel:
// 		return SendLog(title, body, "warning")
// 	case logrus.InfoLevel:
// 		return SendLog(title, body, "info")
// 	case logrus.DebugLevel:
// 		return SendLog(title, body, "debug")
// 	default:
// 		return nil
// 	}
// }

// func (hook *OctokitHook) Levels() []logrus.Level {
// 	return []logrus.Level{
// 		logrus.PanicLevel,
// 		logrus.FatalLevel,
// 		logrus.ErrorLevel,
// 		logrus.WarnLevel,
// 		logrus.InfoLevel,
// 		logrus.DebugLevel,
// 	}
// }
