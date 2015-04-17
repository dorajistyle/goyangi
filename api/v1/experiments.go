package v1

import (
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
)

// @Title Experiments
// @Description Experiments's router group. It contains API in experiments.
func Experiments(parentRoute *gin.RouterGroup) {

	route := parentRoute.Group("/experiments")
	route.GET("/", test)
}

// @Title test
// @Description test API.
// @Accept  json
// @Param   theString        form   string     true        "The string that required."
// @Param   theInt        form   int  false        "The int that optional."
// @Success 200 {object} response.BasicResponse "OK"
// @Failure 401 {object} response.BasicResponse "Authentication required"
// @Resource /experiments/test
// @Router /experiments [get]
func test(c *gin.Context) {
	currentUser, err := userService.CurrentUser(c)
	log.Debugf("header : %s", c.Request.Header.Get("X-Auth-Token"))
	// userService.SetCookieHandler(c, "elsa", "qqqq11")
	if err != nil {
		c.JSON(200, gin.H{"user": nil})
	} else {
		c.JSON(200, gin.H{"user": currentUser})
	}

}
