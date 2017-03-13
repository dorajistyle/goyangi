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
MIICXgIBAAKBgQDXQzg6KtE1y5mq//PfUjelrCRG45FZRW1HKByDbBrzKUP89WYp
izurrPnZtmCaP/zESIgcIv26oHs3b8evq8TpUgHCrt1122jJXBEGIfGrSvL13aon
fLreZBee/EhPxWR0er21SOTGLdzeNIZ9h78pmU7BHBMQcI7BdgchcbtMPwIDAQAB
AoGBAM/m3CdRsz2BpqjBC4hkn9oip+bPY1IU+7X9x4wmnOC8ui1V1ZXKI5drOORM
gIB5zGbGpq6GHQhidp7UFshT3Zi1W8fNRPzRVG7oGH36/HPxDyV5N6UFyXStRq/N
ROpoiuM8mVsiNedqLKMMHMivSNVXxdK7kj5WEwxpC1xrNUwBAkEA7i1VjCtvcY0f
YAliAbwXX2RpGRWS8TNUJsFQ1hMG8DDj2DOdcVCANIwQSBenMEL0wSN2E/zgyUYy
BBqU3BWVqQJBAOde684AJ/kwkKGphonXbSr2QasYeFdQK0e1s8uoPxJVKuCtM0xA
G+kvnlMyLtGXG5q2YYuwvh2BoEkwWFXmM6cCQAPup6zqwqpDRDNXtFCHBHPEup95
ZbWpvUfuhSEjq0en5vsYzw6h35v+e/5UtaPsVxIhPb/SuvtXt1euAKspiBECQQCN
WmGIBnJlfHUwPyjx98o7UB3IkPecqF74vZrt1olKAvxiLY7Ei/pBWZVJ0MPnyoDT
4Y7wz/cmgbZSYJXnTO/LAkEAz7KMzHiKP8vhwP3PCeVK29GlpkLrc86u05zINfbc
Knb2oY2ZhkNA80CzqUMCK50PKfX7rbVpPThHT7mX6rQ+tA==
-----END RSA PRIVATE KEY-----`
		publicKeyRSA = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDXQzg6KtE1y5mq//PfUjelrCRG
45FZRW1HKByDbBrzKUP89WYpizurrPnZtmCaP/zESIgcIv26oHs3b8evq8TpUgHC
rt1122jJXBEGIfGrSvL13aonfLreZBee/EhPxWR0er21SOTGLdzeNIZ9h78pmU7B
HBMQcI7BdgchcbtMPwIDAQAB
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
