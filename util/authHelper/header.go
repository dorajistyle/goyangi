package authHelper

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dorajistyle/goyangi/util/jwt"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
)

// GetTokenString extract a token from an authentication header.
func GetTokenString(c *gin.Context) (string, int, error) {
	var token string
	if c.Request == nil || c.Request.Header == nil {
		return token, http.StatusBadRequest, errors.New("Request header is empty.")
	}
	authorization := c.Request.Header.Get("Authorization")
	if len(authorization) == 0 {
		return token, http.StatusUnauthorized, errors.New("Authorization header is empty.")
	}
	tokens := strings.Split(authorization, " ")
	if len(tokens) < 2 {
		return token, http.StatusUnauthorized, errors.New("Authorization header is not valid.")
	}
	token = tokens[1]
	log.Debugf("token : %s\n", token)
	return token, http.StatusOK, nil
}

// AuthenticateClient authenticate a token that generated from client that is valid.
func AuthenticateClient(c *gin.Context) (string, map[string]string, int, error) {
	log.Debug("AuthenticateClient executed")
	token, status, err := GetTokenString(c)
	if err != nil {
		return token, nil, status, err
	}
	claims, status, err := jwt.ValidateTokenClient(token)
	if err != nil {
		return token, nil, status, err
	}
	claimArray := make(map[string]string)
	claimArray["ak"] = claims["ak"].(string)
	claimArray["sk"] = claims["sk"].(string)
	claimArray["un"] = claims["un"].(string)
	return token, claimArray, status, err
}

// AuthenticateServer authenticate a token that generated from API server that is valid.
func AuthenticateServer(c *gin.Context) (string, map[string]string, int, error) {
	log.Debug("AuthenticateServer executed")
	token, status, err := GetTokenString(c)
	if err != nil {
		return token, nil, status, err
	}
	claims, status, err := jwt.ValidateTokenServer(token)
	if err != nil {
		return token, nil, status, err
	}
	claimArray := make(map[string]string)
	claimArray["ak"] = claims["ak"].(string)
	claimArray["sk"] = claims["sk"].(string)
	claimArray["un"] = claims["un"].(string)
	return token, claimArray, status, err
}
