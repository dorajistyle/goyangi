package userService

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"time"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/crypto"
	"github.com/dorajistyle/goyangi/util/email"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/timeHelper"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nicksnyder/go-i18n/i18n"
)

// SendEmailPasswordResetToken sends a password reset token via email.
func SendEmailPasswordResetToken(to string, token string, locale string) error {
	T, _ := i18n.Tfunc(locale)
	err := email.SendEmailFromAdmin(to,
		T("reset_password_title"),
		T("link")+" : "+config.HostURL+"/reset/password/"+token,
		T("reset_password_content")+"<br/><a href=\""+config.HostURL+"/reset/password/"+token+"\" target=\"_blank\">"+config.HostURL+"/reset/password/"+token+"</a>")
	return err
}

// SendPasswordResetToken sends a password reset token.
func SendPasswordResetToken(c *gin.Context) (int, error) {
	var user model.User
	var sendPasswordResetForm SendPasswordResetForm
	var err error
	log.Debugf("c.Params : %v", c.Params)
	c.BindWith(&sendPasswordResetForm, binding.Form)
	if db.ORM.Where(&model.User{Email: sendPasswordResetForm.Email}).First(&user).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found. Please Check the email.")
	}
	user.PasswordResetUntil = timeHelper.TwentyFourHoursLater()
	user.PasswordResetToken, err = crypto.GenerateRandomToken16()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	log.Debugf("generated token : %s", user.PasswordResetToken)
	status, err := UpdateUserCore(&user)
	if err != nil {
		return status, err
	}
	err = SendEmailPasswordResetToken(user.Email, user.PasswordResetToken, "en-us")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// ResetPassword resets a password of user.
func ResetPassword(c *gin.Context) (int, error) {
	var user model.User
	var passwordResetForm PasswordResetForm
	c.BindWith(&passwordResetForm, binding.Form)
	if db.ORM.Where(&model.User{PasswordResetToken: passwordResetForm.PasswordResetToken}).First(&user).RecordNotFound() {
		return http.StatusNotFound, errors.New("User is not found.")
	}
	isExpired := timeHelper.IsExpired(user.PasswordResetUntil)
	log.Debugf("passwordResetUntil : %s", user.PasswordResetUntil.UTC())
	log.Debugf("expired : %t", isExpired)
	if isExpired {
		return http.StatusForbidden, errors.New("token not valid.")
	}
	newPassword, err := bcrypt.GenerateFromPassword([]byte(passwordResetForm.Password), 10)
	if err != nil {
		return http.StatusInternalServerError, errors.New("User is not updated. Password not Generated.")
	}
	passwordResetForm.Password = string(newPassword)
	log.Debugf("user password before : %s ", user.Password)
	modelHelper.AssignValue(&user, &passwordResetForm)
	user.PasswordResetToken = ""
	user.PasswordResetUntil = time.Now()
	log.Debugf("user password after : %s ", user.Password)
	status, err := UpdateUserCore(&user)
	if err != nil {
		return status, err
	}
	status, err = SetCookie(c, user.Token)
	return status, err
}
