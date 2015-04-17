package v1

import (
	"fmt"

	"github.com/dorajistyle/goyangi/api/response"
	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/service/oauthService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
)

// @Title Oauth
// @Description Oauth's router group.
func Oauth(parentRoute *gin.RouterGroup) {

	route := parentRoute.Group("/oauth")
	route.GET("", retrieveOauthStatus)

	route.GET("/google", googleAuth)
	route.DELETE("/google", userPermission.AuthRequired(googleRevoke))
	route.GET("/google/redirect", googleRedirect)

	route.GET("/github", githubAuth)
	route.DELETE("/github", userPermission.AuthRequired(githubRevoke))
	route.GET("/github/redirect", githubRedirect)

	// route.GET("/yahoo", yahooAuth)
	// route.DELETE("/yahoo", userPermission.AuthRequired(yahooRevoke))
	// route.GET("/yahoo/redirect", yahooRedirect)

	route.GET("/facebook", facebookAuth)
	route.DELETE("/facebook", userPermission.AuthRequired(facebookRevoke))
	route.GET("/facebook/redirect", facebookRedirect)

	// route.GET("/twitter", twitterAuth)
	// route.DELETE("/twitter", userPermission.AuthRequired(twitterRevoke))
	// route.GET("/twitter/redirect", twitterRedirect)

	route.GET("/linkedin", linkedinAuth)
	route.DELETE("/linkedin", userPermission.AuthRequired(linkedinRevoke))
	route.GET("/linkedin/redirect", linkedinRedirect)

	// route.GET("/kakao", kakaoAuth)
	// route.DELETE("/kakao", userPermission.AuthRequired(kakaoRevoke))
	// route.GET("/kakao/redirect", kakaoRedirect)

	// route.GET("/naver", naverAuth)
	// route.DELETE("/naver", userPermission.AuthRequired(naverRevoke))
	// route.GET("/naver/redirect", naverRedirect)
}

// @Title retrieveOauthStatus
// @Description Retrieve oauth connections.
// @Accept  json
// @Success 200 {array} oauthService.oauthStatusMap "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Resource /oauth
// @Router /oauth [get]
func retrieveOauthStatus(c *gin.Context) {
	oauthStatus, status, err := oauthService.RetrieveOauthStatus(c)
	if err == nil {
		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized: "oauth.error.unauthorized",
		}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title googleAuth
// @Description Get google oauth url.
// @Accept  json
// @Success 200 {object} gin.H "OauthURL retrieved"
// @Resource /oauth
// @Router /oauth/google [get]
func googleAuth(c *gin.Context) {
	url, status := oauthService.GoogleURL()
	c.JSON(status, gin.H{"url": url})
}

// @Title googleRevoke
// @Description Get google oauth url.
// @Accept  json
// @Success 200 {object} gin.H "Revoked"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Connection is not found"
// @Failure 500 {object} response.BasicResponse "Connection not revoked from user"
// @Resource /oauth
// @Router /oauth/google [delete]
func googleRevoke(c *gin.Context) {
	oauthStatus, status, err := oauthService.RevokeGoogle(c)
	if err == nil {
		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized:        "oauth.error.unauthorized",
			NotFound:            "oauth.error.notFound",
			InternalServerError: "oauth.error.internalServerError"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title googleRedirect
// @Description Redirect from Google oauth.
// @Accept  json
// @Success 303 {object} response.BasicResponse "Connection linked."
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "User is not found"
// @Failure 500 {object} response.BasicResponse "Connection not linked"
// @Resource /oauth
// @Router /oauth/google/redirect [get]
func googleRedirect(c *gin.Context) {
	status, err := oauthService.OauthGoogle(c)
	if err != nil {
		log.CheckErrorWithMessage(err, fmt.Sprintf("httpStatusCode : %d", status))
	}
	c.Redirect(303, config.HostURL)
}

// @Title githubAuth
// @Description Get github oauth url.
// @Accept  json
// @Success 200 {object} gin.H "{url: oauthURL}"
// @Resource /oauth
// @Router /oauth/github [get]
func githubAuth(c *gin.Context) {
	url, status := oauthService.GithubURL()
	c.JSON(status, gin.H{"url": url})
}

// @Title githubRevoke
// @Description Get github oauth url.
// @Accept  json
// @Success 200 {object} gin.H "Revoked"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Connection is not found"
// @Failure 500 {object} response.BasicResponse "Connection not revoked from user"
// @Resource /oauth
// @Router /oauth/github [delete]
func githubRevoke(c *gin.Context) {
	oauthStatus, status, err := oauthService.RevokeGithub(c)
	if err == nil {
		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized:        "oauth.error.unauthorized",
			NotFound:            "oauth.error.notFound",
			InternalServerError: "oauth.error.internalServerError"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title githubRedirect
// @Description Redirect from Github oauth.
// @Accept  json
// @Success 303 {object} response.BasicResponse "Connection linked."
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "User is not found"
// @Failure 500 {object} response.BasicResponse "Connection not linked"
// @Resource /oauth
// @Router /oauth/github/redirect [get]
func githubRedirect(c *gin.Context) {
	status, err := oauthService.OauthGithub(c)
	if err != nil {
		log.CheckErrorWithMessage(err, fmt.Sprintf("httpStatusCode : %d", status))
	}
	c.Redirect(303, config.HostURL)
}

// TODO: Yahoo Ouath
// // @Title yahooAuth
// // @Description Get yahoo oauth url.
// // @Accept  json
// // @Success 200 {object} gin.H "{url: oauthURL}"
// // @Resource /oauth
// // @Router /oauth/yahoo [get]
// func yahooAuth(c *gin.Context) {
// 	url, status := oauthService.YahooURL()
// 	c.JSON(status, gin.H{"url": url})
// }
//
// // @Title yahooRevoke
// // @Description Get yahoo oauth url.
// // @Accept  json
// // @Success 200 {object} gin.H "Revoked"
// // @Failure 401 {object} response.BasicResponse "Authentication required"
// // @Failure 404 {object} response.BasicResponse "Connection is not found"
// // @Failure 500 {object} response.BasicResponse "Connection not revoked from user"
// // @Resource /oauth
// // @Router /oauth/yahoo [delete]
// func yahooRevoke(c *gin.Context) {
// 	oauthStatus, status, err := oauthService.RevokeYahoo(c)
// 	if err == nil {
// 		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
// } else {
// 	messageTypes := &response.MessageTypes{
// 		Unauthorized:        "oauth.error.unauthorized",
// 		NotFound:            "oauth.error.notFound",
// 		InternalServerError: "oauth.error.internalServerError"}
// 		response.ErrorJSON(c, status, messageTypes, err)
// 	}
// }
// // @Title yahooRedirect
// // @Description Redirect from Yahoo oauth.
// // @Accept  json
// // @Success 303 {object} response.BasicResponse "Connection linked."
// // @Failure 401 {object} response.BasicResponse "Authentication required"
// // @Failure 404 {object} response.BasicResponse "User is not found"
// // @Failure 500 {object} response.BasicResponse "Connection not linked"
// // @Resource /oauth
// // @Router /oauth/yahoo/redirect [get]
// func yahooRedirect(c *gin.Context) {
// 	status, err := oauthService.OauthYahoo(c)
// if err != nil {
// 	log.CheckErrorWithMessage(err, fmt.Sprintf("httpStatusCode : %d", status))
// }
// c.Redirect(303, config.HostURL)
//
// }

// @Title facebookAuth
// @Description Get facebook oauth url.
// @Accept  json
// @Success 200 {object} gin.H "{url: oauthURL}"
// @Resource /oauth
// @Router /oauth/facebook [get]
func facebookAuth(c *gin.Context) {
	url, status := oauthService.FacebookURL()
	c.JSON(status, gin.H{"url": url})
}

// @Title facebookRevoke
// @Description Get facebook oauth url.
// @Accept  json
// @Success 200 {object} gin.H "Revoked"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Connection is not found"
// @Failure 500 {object} response.BasicResponse "Connection not revoked from user"
// @Resource /oauth
// @Router /oauth/facebook [delete]
func facebookRevoke(c *gin.Context) {
	oauthStatus, status, err := oauthService.RevokeFacebook(c)
	if err == nil {
		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized:        "oauth.error.unauthorized",
			NotFound:            "oauth.error.notFound",
			InternalServerError: "oauth.error.internalServerError"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title facebookRedirect
// @Description Redirect from Facebook oauth.
// @Accept  json
// @Success 303 {object} response.BasicResponse "Connection linked."
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "User is not found"
// @Failure 500 {object} response.BasicResponse "Connection not linked"
// @Resource /oauth
// @Router /oauth/facebook/redirect [get]
func facebookRedirect(c *gin.Context) {
	status, err := oauthService.OauthFacebook(c)
	if err != nil {
		log.CheckErrorWithMessage(err, fmt.Sprintf("httpStatusCode : %d", status))
	}
	c.Redirect(303, config.HostURL)
}

// TODO: Twitter  Oauth
// // @Title twitterAuth
// // @Description Get twitter oauth url.
// // @Accept  json
// // @Success 200 {object} gin.H "{url: oauthURL}"
// // @Resource /oauth
// // @Router /oauth/twitter [get]
// func twitterAuth(c *gin.Context) {
// 	url, status := oauthService.TwitterURL()
// 	c.JSON(status, gin.H{"url": url})
// 	// c.JSON(200, response.BasicResponse{})
// }
//
// // @Title twitterRevoke
// // @Description Get twitter oauth url.
// // @Accept  json
// // @Success 200 {object} gin.H "Revoked"
// // @Failure 401 {object} response.BasicResponse "Authentication required"
// // @Failure 404 {object} response.BasicResponse "Connection is not found"
// // @Failure 500 {object} response.BasicResponse "Connection not revoked from user"
// // @Resource /oauth
// // @Router /oauth/twitter [delete]
// func twitterRevoke(c *gin.Context) {
// 	oauthStatus, status, err := oauthService.RevokeTwitter(c)
// 	if err == nil {
// 		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
// } else {
// 	messageTypes := &response.MessageTypes{
// 		Unauthorized:        "oauth.error.unauthorized",
// 		NotFound:            "oauth.error.notFound",
// 		InternalServerError: "oauth.error.internalServerError"}
// 		response.ErrorJSON(c, status, messageTypes, err)
// 	}
// }
// // @Title twitterRedirect
// // @Description Redirect from Twitter oauth.
// // @Accept  json
// // @Success 303 {object} response.BasicResponse "Connection linked."
// // @Failure 401 {object} response.BasicResponse "Authentication required"
// // @Failure 404 {object} response.BasicResponse "User is not found"
// // @Failure 500 {object} response.BasicResponse "Connection not linked"
// // @Resource /oauth
// // @Router /oauth/twitter/redirect [get]
// func twitterRedirect(c *gin.Context) {
// 	status, err := oauthService.OauthTwitter(c)
// if err != nil {
// 	log.CheckErrorWithMessage(err, fmt.Sprintf("httpStatusCode : %d", status))
// }
// c.Redirect(303, config.HostURL)
// }

// @Title linkedinAuth
// @Description Get linkedin oauth url.
// @Accept  json
// @Success 200 {object} gin.H "{url: oauthURL}"
// @Resource /oauth
// @Router /oauth/linkedin [get]
func linkedinAuth(c *gin.Context) {
	url, status := oauthService.LinkedinURL()
	c.JSON(status, gin.H{"url": url})
}

// @Title linkedinRevoke
// @Description Get linkedin oauth url.
// @Accept  json
// @Success 200 {object} gin.H "Revoked"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "Connection is not found"
// @Failure 500 {object} response.BasicResponse "Connection not revoked from user"
// @Resource /oauth
// @Router /oauth/linkedin [delete]
func linkedinRevoke(c *gin.Context) {
	oauthStatus, status, err := oauthService.RevokeLinkedin(c)
	if err == nil {
		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
	} else {
		messageTypes := &response.MessageTypes{
			Unauthorized:        "oauth.error.unauthorized",
			NotFound:            "oauth.error.notFound",
			InternalServerError: "oauth.error.internalServerError"}
		response.ErrorJSON(c, status, messageTypes, err)
	}
}

// @Title linkedinRedirect
// @Description Redirect from Linkedin oauth.
// @Accept  json
// @Success 303 {object} response.BasicResponse "Connection linked."
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Failure 404 {object} response.BasicResponse "User is not found"
// @Failure 500 {object} response.BasicResponse "Connection not linked"
// @Resource /oauth
// @Router /oauth/linkedin/redirect [get]
func linkedinRedirect(c *gin.Context) {
	status, err := oauthService.OauthLinkedin(c)
	if err != nil {
		log.CheckErrorWithMessage(err, fmt.Sprintf("httpStatusCode : %d", status))
	}
	c.Redirect(303, config.HostURL)
}

// // @Title kakaoAuth
// // @Description Get kakao oauth url.
// // @Accept  json
// // @Success 200 {object} gin.H "{url: oauthURL}"
// // @Resource /oauth
// // @Router /oauth/kakao [get]
// func kakaoAuth(c *gin.Context) {
// 	url, status := oauthService.KakaoURL()
// 	c.JSON(status, gin.H{"url": url})
// }
//
// // @Title kakaoRevoke
// // @Description Get kakao oauth url.
// // @Accept  json
// // @Success 200 {object} gin.H "Revoked"
// // @Failure 401 {object} response.BasicResponse "Authentication required"
// // @Failure 404 {object} response.BasicResponse "Connection is not found"
// // @Failure 500 {object} response.BasicResponse "Connection not revoked from user"
// // @Resource /oauth
// // @Router /oauth/kakao [delete]
// func kakaoRevoke(c *gin.Context) {
// 	oauthStatus, status, err := oauthService.RevokeKakao(c)
// 	if err == nil {
// 		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
// } else {
// 	messageTypes := &response.MessageTypes{
// 		Unauthorized:        "oauth.error.unauthorized",
// 		NotFound:            "oauth.error.notFound",
// 		InternalServerError: "oauth.error.internalServerError"}
// 		response.ErrorJSON(c, status, messageTypes, err)
// 	}
// }
// // @Title kakaoRedirect
// // @Description Redirect from Kakao oauth.
// // @Accept  json
// // @Success 303 {object} response.BasicResponse "Connection linked."
// // @Failure 401 {object} response.BasicResponse "Authentication required"
// // @Failure 404 {object} response.BasicResponse "User is not found"
// // @Failure 500 {object} response.BasicResponse "Connection not linked"
// // @Resource /oauth
// // @Router /oauth/kakao/redirect [get]
// func kakaoRedirect(c *gin.Context) {
// 	status, err := oauthService.OauthKakao(c)
// if err != nil {
// 	log.CheckErrorWithMessage(err, fmt.Sprintf("httpStatusCode : %d", status))
// }
// c.Redirect(303, config.HostURL)
//
// }

// // @Title naverAuth
// // @Description Get naver oauth url.
// // @Accept  json
// // @Success 200 {object} gin.H "{url: oauthURL}"
// // @Resource /oauth
// // @Router /oauth/naver [get]
// func naverAuth(c *gin.Context) {
// 	url, status := oauthService.NaverURL()
// 	c.JSON(status, gin.H{"url": url})
// 	// c.JSON(200, response.BasicResponse{})
// }
//
// // @Title naverRevoke
// // @Description Get naver oauth url.
// // @Accept  json
// // @Success 200 {object} gin.H "Revoked"
// // @Failure 401 {object} response.BasicResponse "Authentication required"
// // @Failure 404 {object} response.BasicResponse "Connection is not found"
// // @Failure 500 {object} response.BasicResponse "Connection not revoked from user"
// // @Resource /oauth
// // @Router /oauth/naver [delete]
// func naverRevoke(c *gin.Context) {
// 	oauthStatus, status, err := oauthService.RevokeNaver(c)
// 	if err == nil {
// 		c.JSON(status, gin.H{"oauthStatus": oauthStatus})
// } else {
// 	messageTypes := &response.MessageTypes{
// 		Unauthorized:        "oauth.error.unauthorized",
// 		NotFound:            "oauth.error.notFound",
// 		InternalServerError: "oauth.error.internalServerError"}
// 		response.ErrorJSON(c, status, messageTypes, err)
// 	}
// }
// // @Title naverRedirect
// // @Description Redirect from Naver oauth.
// // @Accept  json
// // @Success 303 {object} response.BasicResponse "Connection linked."
// // @Failure 401 {object} response.BasicResponse "Authentication required"
// // @Failure 404 {object} response.BasicResponse "User is not found"
// // @Failure 500 {object} response.BasicResponse "Connection not linked"
// // @Resource /oauth
// // @Router /oauth/naver/redirect [get]
// func naverRedirect(c *gin.Context) {
// 	status, err := oauthService.OauthNaver(c)
// if err != nil {
// 	log.CheckErrorWithMessage(err, fmt.Sprintf("httpStatusCode : %d", status))
// }
// c.Redirect(303, config.HostURL)
//
// }
