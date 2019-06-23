package jwt_test

import (
	"time"

	. "github.com/dorajistyle/goyangi/util/jwt"
	jwt "github.com/dorajistyle/goyangi/vendor/github.com/dgrijalva/jwt-go"

	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jwt", func() {
	var (
		appkey            string
		secretkey         string
		userName          string
		expiration        int64
		expirationExpired int64
		userToken         string
		token             *jwt.Token
		claims            map[string]interface{}
		status            int
		err               error
		hmacKey           string
		privateKeyRSA     string
		publicKeyRSA      string
	)

	BeforeEach(func() {
		hmacKey = "secret"
		privateKeyRSA = `-----BEGIN RSA PRIVATE KEY-----
blablabla==
-----END RSA PRIVATE KEY-----`
		publicKeyRSA = `-----BEGIN PUBLIC KEY-----
blablabla
-----END PUBLIC KEY-----`
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

	Describe("Parse a user token by a RSA method", func() {
		Context("when parse the token and err", func() {
			BeforeEach(func() {
				userToken, err = CreateTokenRSA(appkey, secretkey, userName, expiration, privateKeyRSA)
				token, err = ParseTokenRSA(userToken, publicKeyRSA)
			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})

		})
	})

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
				token, err = ParseTokenHMAC(userToken, hmacKey)
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
