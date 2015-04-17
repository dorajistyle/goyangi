package userService

import (
	"errors"
	"net/http"

	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

// // UserName get username from a cookie.
// func UserName(c *gin.Context) (string, error) {
// 	var userName string
// 	request := c.Request
// 	cookie, err := request.Cookie("session")
// 	if err != nil {
// 		return userName, err
// 	}
// 	cookieValue := make(map[string]string)
// 	err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
// 	if err != nil {
// 		return userName, err
// 	}
// 	userName = cookieValue["name"]
// 	return userName, nil
// }

func Token(c *gin.Context) (string, error) {
	var token string
	request := c.Request
	cookie, err := request.Cookie("session")
	if err != nil {
		return token, err
	}
	cookieValue := make(map[string]string)
	err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
	if err != nil {
		return token, err
	}
	token = cookieValue["token"]
	if len(token) == 0 {
		return token, errors.New("Token is empty.")
	}
	return token, nil
}

// SetCookie sets a cookie.
func SetCookie(c *gin.Context, token string) (int, error) {
	response := c.Writer
	value := map[string]string{
		"token": token,
	}
	encoded, err := cookieHandler.Encode("session", value)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}
	http.SetCookie(response, cookie)
	return http.StatusOK, nil
}

// ClearCookie clears a cookie.
func ClearCookie(c *gin.Context) (int, error) {
	response := c.Writer
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
	return http.StatusOK, nil
}

// SetCookieHandler sets a cookie with email and password.
func SetCookieHandler(c *gin.Context, email string, pass string) (int, error) {
	if email != "" && pass != "" {
		log.Debugf("User email : %s , password : %s", email, pass)
		var user model.User
		isValidEmail := validation.EmailValidation(email)
		if isValidEmail {
			log.Debug("User entered valid email.")
			if db.ORM.Where("email = ?", email).First(&user).RecordNotFound() {
				return http.StatusNotFound, errors.New("User is not found.")
			}
		} else {
			log.Debug("User entered username.")
			if db.ORM.Where("username = ?", email).First(&user).RecordNotFound() {
				return http.StatusNotFound, errors.New("User is not found.")
			}
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
		if err != nil {
			return http.StatusUnauthorized, errors.New("Password incorrect.")
		}
		status, err := SetCookie(c, user.Token)
		if err != nil {
			return status, err
		}
		c.Writer.Header().Set("X-Auth-Token", user.Token)
		return http.StatusOK, nil
	} else {
		return http.StatusNotFound, errors.New("User is not found.")
	}
}

// RegisterHanderFromForm sets cookie from a RegistrationForm.
func RegisterHanderFromForm(c *gin.Context, registrationForm RegistrationForm) (int, error) {
	email := registrationForm.Email
	pass := registrationForm.Password
	log.Debugf("RegisterHandler UserEmail : %s", email)
	log.Debugf("RegisterHandler UserPassword : %s", pass)
	status, err := SetCookieHandler(c, email, pass)
	return status, err
}

// RegisterHandler sets a cookie when user registered.
func RegisterHandler(c *gin.Context) (int, error) {
	var registrationForm RegistrationForm
	c.BindWith(&registrationForm, binding.Form)
	status, err := RegisterHanderFromForm(c, registrationForm)
	return status, err
}

// CurrentUser get a current user.
func CurrentUser(c *gin.Context) (model.User, error) {
	var user model.User
	var token string
	var err error
	token = c.Request.Header.Get("X-Auth-Token")
	if len(token) > 0 {
		log.Debug("header token exist.")
	} else {
		token, err = Token(c)
		log.Debug("header token not exist.")
		if err != nil {
			return user, err
		}
	}
	if db.ORM.Select(config.UserPublicFields+", email").Where("token = ?", token).First(&user).RecordNotFound() {
		return user, errors.New("User is not found.")
	}
	db.ORM.Model(&user).Association("Languages").Find(&user.Languages)
	db.ORM.Model(&user).Association("Roles").Find(&user.Roles)
	return user, nil
}
