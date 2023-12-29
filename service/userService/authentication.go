package userService

import (
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// CreateUserAuthentication creates user authentication.
func CreateUserAuthentication(c *gin.Context) (int, error) {
	var form LoginForm
	bindErr := c.MustBindWith(&form, binding.Form)
	log.Debugf("bind error : %s\n", bindErr)
	email := form.Email
	pass := form.Password
	status, err := SetCookieHandler(c, email, pass)
	return status, err
}
