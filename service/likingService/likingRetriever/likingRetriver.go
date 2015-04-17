package likingRetriever

import (
	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/pagination"
)

// RetrieveLikings retrieves likings.
func RetrieveLikings(item interface{}, currentPages ...int) ([]*model.PublicUser, int, bool, bool, int) {
	var users []*model.User
	var currentPage int
	if len(currentPages) > 0 {
		currentPage = currentPages[0]
	} else {
		currentPage = 1
	}
	count := db.ORM.Model(item).Association("Likings").Count()
	offset, currentPage, hasPrev, hasNext := pagination.Paginate(currentPage, config.LikingPerPage, count)
	db.ORM.Limit(config.LikingPerPage).Order(config.LikingOrder).Offset(offset).Select(config.UserPublicFields).Model(item).Association("Likings").Find(&users)
	var likingArr []*model.PublicUser
	for _, user := range users {
		likingArr = append(likingArr, &model.PublicUser{User: user})
	}

	log.Debugf("likings : %v", likingArr)
	return likingArr, currentPage, hasPrev, hasNext, count
}

// RetrieveLiked retrieves liked.
func RetrieveLiked(item interface{}, currentPages ...int) ([]*model.PublicUser, int, bool, bool, int) {
	var users []*model.User
	var currentPage int
	if len(currentPages) > 0 {
		currentPage = currentPages[0]
	} else {
		currentPage = 1
	}
	count := db.ORM.Model(item).Association("Liked").Count()
	offset, currentPage, hasPrev, hasNext := pagination.Paginate(currentPage, config.LikedPerPage, count)
	db.ORM.Limit(config.LikedPerPage).Order(config.LikedOrder).Offset(offset).Select(config.UserPublicFields).Model(item).Association("Liked").Find(&users)

	var likedArr []*model.PublicUser
	for _, user := range users {
		likedArr = append(likedArr, &model.PublicUser{User: user})
	}

	log.Debugf("liked : %v", users)
	return likedArr, currentPage, hasPrev, hasNext, count
}
