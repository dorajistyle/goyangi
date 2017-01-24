package v1

import (
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
)

// @Title Experiments
// @Description Experiments's router group. It contains API in experiments.
func Experiments(parentRoute *gin.RouterGroup) {

	route := parentRoute.Group("/experiments")
	route.GET("", test)
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
	var user model.User

	// var users []model.User

	db.ORM.First(&user)

	log.Debugf("user : %v", user)
	log.Debugf("user : %s", user.Name)
	var roles []model.Role

	err = db.ORM.Model(user).Association("Roles").Find(&roles).Error
	if err != nil {
		log.Fatalf("fatal : %v", err)
	}
	log.Debugf("roles : %v", roles)

	var likings []model.User
	if err = db.ORM.Joins("JOIN users_followers on users_followers.user_id=?", user.Id).
		Group("users.id").Find(&likings).Error; err != nil {
		log.Fatal(err.Error())
	}
	user.Likings = likings
	log.Debugf("Likings : %v", likings)
	var liked []model.User
	if err = db.ORM.Joins("JOIN users_followers on users_followers.follower_id=?", user.Id).
		Group("users.id").Find(&liked).Error; err != nil {
		log.Fatal(err.Error())
	}
	user.Liked = liked
	log.Debugf("Liked : %v", liked)
	// var follwerIDs []uint
	// rows, _ := db.ORM.Table("users_followers").Where("user_id = ?", user.Id).Select("follower_id").Rows()
	// for rows.Next() {
	// 	var followerID uint
	// 	rows.Scan(&followerID)
	// 	follwerIDs = append(follwerIDs, followerID)
	// }
	// log.Debugf("follwerIDs : %v", follwerIDs)

	// user.Likings =
	// err = db.ORM.Model(user).Association("Likings").Find(&users).Error
	// if err != nil {
	// 	log.Fatalf("fatal : %v", err)
	// }

	// log.Debugf("users : %v", users[0].User)
	// log.Debugf("users : %v", users[0].Follower)
	// var liked []model.User
	// err = db.ORM.Model(user).Association("Liked").Find(&liked).Error
	// if err != nil {
	// 	log.Fatalf("fatal : %v", err)
	// }
	// log.Debugf("liked : %v", liked)

	// userService.SetCookieHandler(c, "elsa", "qqqq11")
	if err != nil {
		c.JSON(200, gin.H{"user": nil})
	} else {
		c.JSON(200, gin.H{"user": currentUser})
	}

}
