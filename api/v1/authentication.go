package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/dorajistyle/goyangi/api/response"
	"github.com/dorajistyle/goyangi/service/userService"
)

// @Title Authentications
// @Description Authentications's router group.
func Authentications(parentRoute *gin.RouterGroup) {
	route := parentRoute.Group("/authentications")
	route.POST("/", createUserAuthentication)
	route.DELETE("/", deleteUserAuthentication)
}

// @Title createUserAuthentication
// @Description Create a user session.
// @Accept  json
// @Param   loginEmail        form   string     true        "User email."
// @Param   loginPassword        form   string  true        "User password."
// @Success 201 {object} response.BasicResponse "User authentication created"
// @Failure 401 {object} response.BasicResponse "Password incorrect"
// @Failure 404 {object} response.BasicResponse "User is not found"
// @Resource /authentications
// @Router /authentications [post]
func createUserAuthentication(c *gin.Context) {
	status, err := userService.CreateUserAuthentication(c)
	messageTypes := &response.MessageTypes{OK: "login.done",
		Unauthorized: "login.error.passwordIncorrect",
		NotFound:     "login.error.userNotFound"}
	messages := &response.Messages{OK: "User logged in successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}

// @Title deleteUserAuthentication
// @Description Delete a user session.
// @Accept  json
// @Success 200 {object}  response.BasicResponse "User logged out successfully"
// @Resource /authentications
// @Router /authentications [delete]
func deleteUserAuthentication(c *gin.Context) {
	status, err := userService.ClearCookie(c)
	messageTypes := &response.MessageTypes{OK: "logout.done"}
	messages := &response.Messages{OK: "User logged out successfully."}
	response.JSON(c, status, messageTypes, messages, err)
}
