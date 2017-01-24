package authHelper_test

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/config"
	"github.com/dorajistyle/goyangi/util/jwt"

	. "github.com/dorajistyle/goyangi/util/authHelper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("authHelper", func() {
	var (
		appKey     string
		secretkey  string
		userName   string
		userToken  string
		expiration int64
		claims     map[string]string
		status     int
		err        error
		c          *gin.Context
		user       model.User
	)

	BeforeEach(func() {
		c, _, _ = gin.CreateTestContext()
		c.Request, _ = http.NewRequest("POST", "/api/v1/test", nil)

		appKey = "TESTugsiEHS4Ycx7uBI88DE6ZFo7jAl4"
		secretkey = "TESTugsiEHS4Ycx7uBI88DE6ZFo7jAl4"
		userName = "John Busker Tester"
		expiration = time.Now().Add(time.Hour * 24).Unix()

	})

	Describe("Authenticate a token from client", func() {

		Context("when authenticate the token", func() {
			BeforeEach(func() {
				userToken, err = jwt.CreateTokenHMAC(appKey, secretkey, userName, expiration, config.JWTSigningKeyHMACClient)
				c.Request.Header.Set("Authorization", "bearer "+userToken)
				_, claims, status, err = AuthenticateClient(c)
			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})

		})
	})

	Describe("Authenticate a token from server", func() {

		Context("when authenticate the token", func() {
			BeforeEach(func() {
				userToken, status, err = jwt.CreateToken(appKey, secretkey, userName)
				c.Request.Header.Set("Authorization", "bearer "+userToken)
				// _, status, err = jwt.ValidateToken(userToken, config.JWTSigningKeyHMACServer)
				_, claims, status, err = AuthenticateServer(c)

			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})

		})
	})

	Describe("Get an authorized user", func() {

		Context("when get the authorized user", func() {
			BeforeEach(func() {
				userToken, status, err = jwt.CreateToken(appKey, secretkey, userName)
				c.Request.Header.Set("Authorization", "bearer "+userToken)
				_, claims, status, err = AuthenticateServer(c)
				user, status, err = GetAuthorizedUser(claims["ak"], claims["sk"], claims["un"])
			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})
			It("UserName should be claims[\"un\"].", func() {
				Expect(user.Name).To(Equal(claims["un"]))
			})

		})
	})

	Describe("Get an authorized user from context", func() {

		Context("when get the authorized user", func() {
			BeforeEach(func() {
				userToken, status, err = jwt.CreateToken(appKey, secretkey, userName)
				c.Request.Header.Set("Authorization", "bearer "+userToken)
				user, status, err = GetAuthorizedUserFromContext(c)
			})
			It("err should be nil.", func() {
				Expect(err).To(BeNil())
			})
		})
	})
})
