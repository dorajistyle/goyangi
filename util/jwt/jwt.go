package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/dorajistyle/goyangi/util/timeHelper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"

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
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingKey), nil
	})
	return token, err
}

// CreateTokenRSA creates a jwt token by RSA method.
func CreateTokenRSA(appKey string, secretkey string, username string, expiration int64, signingKeyString string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	// Set some claims
	token.Claims.(jwt.MapClaims)["ak"] = appKey
	token.Claims.(jwt.MapClaims)["sk"] = secretkey
	token.Claims.(jwt.MapClaims)["un"] = username
	token.Claims.(jwt.MapClaims)["exp"] = expiration

	block, _ := pem.Decode([]byte(signingKeyString))
	signingKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

// ParseTokenRSA parses token by RSA method.
// TODO: Error should be fixed. : token signature is invalid: key is of invalid type
func ParseTokenRSA(userToken string, publicKey string) (*jwt.Token, error) {
	// _, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	// return nil, err
	token, err := jwt.Parse(userToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(publicKey), nil
	})
	return token, err
}

// GenerateRSAKeys generate RSA keys(private and public).
func GenerateRSAKeys(bits int) {
	signingKey, _ := rsa.GenerateKey(rand.Reader, bits) // bits could be 512, 1024, 2048, 4096

	signingBytes := x509.MarshalPKCS1PrivateKey(signingKey)

	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: signingBytes,
	}

	generatedPrivateKey := pem.EncodeToMemory(&block)

	fmt.Printf("%s\n", generatedPrivateKey)
	publicKey := signingKey.PublicKey
	publicBytes, _ := x509.MarshalPKIXPublicKey(&publicKey)
	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicBytes,
	}
	generatedPublicKey := pem.EncodeToMemory(&block)

	fmt.Printf("%s\n", generatedPublicKey)
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
	expiration := timeHelper.FewDurationLaterMillisecond(time.Hour * viper.GetDuration("jwt.server.expirationHour"))

	switch viper.GetString("jwt.server.method") {
	case "HMAC256":
		token, err = CreateTokenHMAC(appKey, secretkey, username, expiration, viper.GetString("jwt.server.key.private"))
	case "RSA256":
		token, err = CreateTokenRSA(appKey, secretkey, username, expiration, viper.GetString("jwt.server.key.private"))
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
	switch viper.GetString("jwt.server.method") {
	case "HMAC256":
		token, err = ParseTokenHMAC(userToken, signingKey)
	case "RSA256":
		token, err = ParseTokenRSA(userToken, signingKey)
	}
	if err != nil {
		if err == jwt.ErrInvalidKey {
			if err == jwt.ErrTokenMalformed {
				return nil, http.StatusUnauthorized, err
			} else if (err == jwt.ErrTokenExpired) || (err == jwt.ErrTokenNotValidYet) {
				claims = token.Claims.(jwt.MapClaims)
				if signingKey == viper.GetString("jwt.server.key.private") {
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
	switch viper.GetString("jwt.server.method") {
	case "HMAC256":
		signingKey = viper.GetString("jwt.server.key.private")
	case "RSA256":
		signingKey = viper.GetString("jwt.server.key.public")
	}
	return ValidateToken(userToken, signingKey)
}

// ValidateTokenClient validate a token that generated from client
func ValidateTokenClient(userToken string) (map[string]interface{}, int, error) {
	var signingKey string
	switch viper.GetString("jwt.client.method") {
	case "HMAC256":
		signingKey = viper.GetString("jwt.client.key.private")
	case "RSA256":
		signingKey = viper.GetString("jwt.client.key.public")
	}
	return ValidateToken(userToken, signingKey)
}
