package userService

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateUserAuthentication creates user authentication.
func CreateUserAuthentication(c *gin.Context) (int, error) {
	var form LoginForm
	c.BindWith(&form, binding.Form)
	email := form.Email
	pass := form.Password
	status, err := SetCookieHandler(c, email, pass)
	return status, err
}
