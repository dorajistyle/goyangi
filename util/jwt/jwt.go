package jwt

import (
	"github.com/dorajistyle/goyangi/util/config"
	"github.com/dorajistyle/goyangi/util/timeHelper"
	"github.com/dgrijalva/jwt-go"
	// "github.com/dorajistyle/goyangi/util/interfaceHelper"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
)

// CreateTokenHMAC creates a jwt token by HMAC method.
func CreateTokenHMAC(appKey string, secretkey string, username string, expiration int64, signingKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims.(jwt.MapClaims)["ak"] = appKey
	token.Claims.(jwt.MapClaims)["sk"] = secretkey
	token.Claims.(jwt.MapClaims)["un"] = username
	token.Claims.(jwt.MapClaims)["exp"] = expiration
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(signingKey))
	return tokenString, err
}

// ParseTokenHMAC parses token by HMAC method.
func ParseTokenHMAC(userToken string, signingKey string) (*jwt.Token, error) {
	// map[string]interface{}
	// token, err := jwt.Parse(userToken, func(token *jwt.Token) ([]byte, error) {
	// 		return []byte(config.JWTSigningKey), nil
	// 	})

	// token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
	//     return []byte(config.JWTSigningKey), nil
	// })
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// return myLookupKey(token.Header["kid"]), nil
		return []byte(signingKey), nil
	})
	return token, err
	// if err == nil && token.Valid {
	//     return token, nil
	//     // return token.Claims, nil
	//     // return token.Claims["ak"].(string), nil
	//     // deliverGoodness("!")
	// } else {
	//     return nil, err
	//     // deliverUtterRejection(":(")
	// }
}

// CreateTokenRSA creates a jwt token by RSA method.
func CreateTokenRSA(appKey string, secretkey string, username string, expiration int64, signingKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	// Set some claims
	token.Claims.(jwt.MapClaims)["ak"] = appKey
	token.Claims.(jwt.MapClaims)["sk"] = secretkey
	token.Claims.(jwt.MapClaims)["un"] = username
	token.Claims.(jwt.MapClaims)["exp"] = expiration
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(signingKey))
	return tokenString, err
}

// ParseTokenRSA parses token by RSA method.
func ParseTokenRSA(userToken string, signingKey string) (*jwt.Token, error) {
	// map[string]interface{}
	// token, err := jwt.Parse(userToken, func(token *jwt.Token) ([]byte, error) {
	// 		return []byte(config.JWTSigningKey), nil
	// 	})

	// token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
	//     return []byte(config.JWTSigningKey), nil
	// })
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// return myLookupKey(token.Header["kid"]), nil
		return []byte(signingKey), nil
	})
	return token, err
	// if err == nil && token.Valid {
	//     return token, nil
	//     // return token.Claims, nil
	//     // return token.Claims["ak"].(string), nil
	//     // deliverGoodness("!")
	// } else {
	//     return nil, err
	//     // deliverUtterRejection(":(")
	// }
}

// CreateToken create a jwt token.
func CreateToken(appKey string, secretkey string, username string) (string, int, error) {
	var token string
	var err error
	if len(appKey) <= 0 {
		return "AppKey required", http.StatusPreconditionFailed, errors.New("AppKey is required to create an authorization token.")
	}
	if len(secretkey) <= 0 {
		return "Secretkey required", http.StatusPreconditionFailed, errors.New("Secretkey is required to create an authorization token.")
	}
	if len(username) <= 0 {
		return "Username required", http.StatusPreconditionFailed, errors.New("Username is required to create an authorization token.")
	}
	expiration := timeHelper.FewDurationLaterMillisecond(time.Hour * config.JWTExpriationHourServer)

	switch config.JWTSigningMethodServer {
	case "HMAC256":
		token, err = CreateTokenHMAC(appKey, secretkey, username, expiration, config.JWTSigningKeyHMACServer)
	case "RSA256":
		token, err = CreateTokenRSA(appKey, secretkey, username, expiration, config.JWTSigningKeyRSAServer)
	}
	if err != nil {
		return "Token creation error", http.StatusUnauthorized, err
	}
	return token, http.StatusOK, nil

}

// ValidateToken validates a token.
func ValidateToken(userToken string, signingKey string) (map[string]interface{}, int, error) {
	var err error
	var token *jwt.Token
	var claims map[string]interface{}
	switch config.JWTSigningMethodServer {
	case "HMAC256":
		token, err = ParseTokenHMAC(userToken, signingKey)
	case "RSA256":
		token, err = ParseTokenRSA(userToken, signingKey)
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, http.StatusUnauthorized, err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				claims = token.Claims.(jwt.MapClaims)
				if signingKey == config.JWTSigningKeyHMACServer {
					var app model.App
					var user model.User
					if db.ORM.First(&app, "key = ? and secret_key = ?", claims["ak"], claims["sk"]).RecordNotFound() {
						return nil, http.StatusNotFound, errors.New("App is not found.")
					}

					if db.ORM.Where("app_id=? and name=?", app.Id, claims["un"]).First(&user).RecordNotFound() {
						return nil, http.StatusNotFound, errors.New("User not found. (Check the AppId and the UserName)")
					}
					var expiredTokenLog model.ExpiredTokenLog

					expiredTokenLog.UserId = user.Id
					expiredTokenLog.AccessedAt = time.Now()
					if db.ORM.Create(&expiredTokenLog).Error != nil {
						return nil, http.StatusBadRequest, errors.New("ExpiredTokenLog is not created.")
					}
				}
			} else {
				return nil, http.StatusUnauthorized, err
			}
		} else {
			return nil, http.StatusUnauthorized, err
		}
	}
	claims = token.Claims.(jwt.MapClaims)
	// var expInt64 int64
	// exp := claims["exp"]
	// expInt64, err = interfaceHelper.GetInt64(exp)
	// if err != nil || exp == nil || timeHelper.IsExpired(time.Unix(expInt64,0)) {
	//     var app model.App
	//     var user model.User
	//     if db.ORM.First(&app, "key = ? and secret_key = ?", claims["ak"], claims["sk"]).RecordNotFound() {
	//         return nil, http.StatusNotFound, errors.New("App is not found.")
	//     }

	//     if db.ORM.Where("app_id=? and name=?", app.Id, claims["un"]).First(&user).RecordNotFound() {
	//         return nil, http.StatusNotFound, errors.New("User not found. (Check the AppId and the UserName)")
	//     }
	//     var expiredTokenLog model.ExpiredTokenLog

	//     expiredTokenLog.UserId = user.Id
	//     expiredTokenLog.AccessedAt = time.Now()
	//     createErr := db.ORM.Create(&expiredTokenLog).Error
	//     if  createErr != nil {
	//         return nil, http.StatusBadRequest, createErr
	//     }
	//     // return nil, http.StatusUnauthorized, errors.New("Token is expired")
	// }
	return claims, http.StatusOK, nil
}

// ValidateTokenServer validate a token that generated from API server
func ValidateTokenServer(userToken string) (map[string]interface{}, int, error) {
	var signingKey string
	switch config.JWTSigningMethodServer {
	case "HMAC256":
		signingKey = config.JWTSigningKeyHMACServer
	case "RSA256":
		signingKey = config.JWTPublicKeyRSAServer
	}
	return ValidateToken(userToken, signingKey)
}

// ValidateTokenClient validate a token that generated from client
func ValidateTokenClient(userToken string) (map[string]interface{}, int, error) {
	var signingKey string
	switch config.JWTSigningMethodClient {
	case "HMAC256":
		signingKey = config.JWTSigningKeyHMACClient
	case "RSA256":
		signingKey = config.JWTPublicKeyRSAClient
	}
	return ValidateToken(userToken, signingKey)
}
