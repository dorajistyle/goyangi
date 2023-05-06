package jwt_test

import (
	"time"

	. "github.com/dorajistyle/goyangi/util/jwt"
	viper "github.com/dorajistyle/goyangi/util/viper"

	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	viper.LoadConfig()
}

var _ = Describe("Jwt", func() {
	var (
		appkey            string
		secretkey         string
		userName          string
		expiration        int64
		expirationExpired int64
		userToken         string
		claims            map[string]interface{}
		status            int
		err               error
		hmacKey           string
		privateKeyRSA     string
		publicKeyRSA      string
	)

	BeforeEach(func() {
		hmacKey = "secret"
		privateKeyRSA = `-----BEGIN PRIVATE KEY-----
MIICWwIBAAKBgQCVCsacV4bK5KU749Ve56PtDDoXVJZ4mzIcA/I+S7CqwCTQXkxq
alGkVCeI72sgnNbEj5n4pbRdxBBD1thXD/dOdYHRSRXeJIMdCKO/v1E9iQJmF0B/
FfkJ+KAaj87e5uWKo5XOPi+25DYsS050B/NFy7bbSjjtk8gCl+0DV/IuBQIDAQAB
AoGAFCfr6jLQENpRGkNalMYg3ir8JDGVU+QxJ6bE+PXFg6IOmHtYPD/6oI2c9yDh
zPxI8zY0bXMDbHbaeEIy6btIB4h4YEqZ3I3Pz+/+vfW1cxMY03UXOSkyD3ZG77LT
N8Bd3HQhJ/1EKeOneT7WWw/MLqzku11cvGHEEGOaFcTEOjUCQQDA5N2vuQaDg3c4
jTfPiJ28pfbZmFHY6NOlQxEmNdwZEB13nFOLbtrAdGIAJqUye89tzlk5huh62ZUO
sOmJbjGrAkEAxc1AsQOsKvBJ4VdRxBH3nn72xH9MZHRzNCDQZvFZ0KWKsQ4R+ylU
TYq2CW0/kG0x4WCDn0sc+FZR0SR95PzPDwJAcIKN07smU3tRBMlJ7mEPMEPVkeHI
i65yFIjj7deog23k4ilqiX+lVHAN4WypGqMgwDmFzYok+9MBoEoMTb7adQJAcAsj
AOImrS/teZKfw2O2EvayS34cRK7d7wJDanx+Nrz+wepJby7rDP1svgw/PE1OOu8T
v7CpmVY0BDcahRJbKwJAVV6QV6adfnKHJsaHfaGJQ2S716TEK9heGGHmmCVHZA8v
HhUJrq0UfaNmj8vf/Af6yzhXN0IwdSpSEzHGpKESXQ==
-----END PRIVATE KEY-----`
		publicKeyRSA = `-----BEGIN RSA PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCVCsacV4bK5KU749Ve56PtDDoX
VJZ4mzIcA/I+S7CqwCTQXkxqalGkVCeI72sgnNbEj5n4pbRdxBBD1thXD/dOdYHR
SRXeJIMdCKO/v1E9iQJmF0B/FfkJ+KAaj87e5uWKo5XOPi+25DYsS050B/NFy7bb
Sjjtk8gCl+0DV/IuBQIDAQAB
-----END RSA PUBLIC KEY-----`
		appkey = "APP-KEY-STRING-FOR-TEST"
		secretkey = "SECRET-KEY-STRING-FOR-TEST"
		userName = "USERNAME-STRING-FOR-TEST"
		expiration = time.Now().Add(time.Hour * 24).Unix()
		expirationExpired = time.Now().Unix()
	})

	Describe("Create a valid token by a RSA method", func() {
		Context("when getting the token and err", func() {
			BeforeEach(func() {
				userToken, err = CreateTokenRSA(appkey, secretkey, userName, expiration, privateKeyRSA)
			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})

		})
	})

	// TODO: Error should be fixed. : token signature is invalid: key is of invalid type
	// Describe("Parse a user token by a RSA method", func() {
	// 	Context("when parse the token and err", func() {
	// 		BeforeEach(func() {
	// 			userToken, err = CreateTokenRSA(appkey, secretkey, userName, expiration, privateKeyRSA)
	// 			_, err = ParseTokenRSA(userToken, publicKeyRSA)
	// 		})
	// 		It("err should be nil.", func() {
	// 			Expect(err).To(BeNil())
	// 		})

	// 	})
	// })

	Describe("Create a valid token by a HMAC method", func() {

		Context("when getting the token and err", func() {
			BeforeEach(func() {
				userToken, err = CreateTokenHMAC(appkey, secretkey, userName, expiration, hmacKey)
			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})

		})
	})

	Describe("Parse a user token by a HMAC method", func() {
		Context("when parse the token and err", func() {
			BeforeEach(func() {
				userToken, err = CreateTokenHMAC(appkey, secretkey, userName, expiration, hmacKey)
				_, err = ParseTokenHMAC(userToken, hmacKey)
			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Validate a user token", func() {
		Context("when parse the token and err", func() {
			BeforeEach(func() {
				userToken, status, err = CreateToken(appkey, secretkey, userName)
				claims, status, err = ValidateTokenServer(userToken)
			})
			It("claim[\"ak\"] should be appkey.", func() {
				Expect(claims["ak"]).To(Equal(appkey))
			})
			It("status should be 200.", func() {
				Expect(status).To(Equal(http.StatusOK))
			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Validate a expired user token by a RSA method", func() {
		Context("when parse the token and err", func() {
			BeforeEach(func() {
				userToken, err = CreateTokenRSA(appkey, secretkey, userName, expirationExpired, privateKeyRSA)
				claims, status, err = ValidateToken(userToken, publicKeyRSA)
			})
			It("status should be 401.", func() {
				Expect(status).To(Equal(http.StatusUnauthorized))
			})
			It("err should be not nil.", func() {
				Expect(err).NotTo(BeNil())
			})
		})
		Context("when parse the token and err", func() {
			BeforeEach(func() {
				userToken, err = CreateTokenRSA(appkey, secretkey, userName, expirationExpired, privateKeyRSA)
				claims, status, err = ValidateToken(userToken, publicKeyRSA)
			})
			It("status should be 401.", func() {
				Expect(status).To(Equal(http.StatusUnauthorized))
			})
			It("err should be not nil.", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
