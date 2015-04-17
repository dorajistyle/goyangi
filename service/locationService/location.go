package locationService

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/commentService"
	"github.com/dorajistyle/goyangi/service/likingService/likingMeta"
	"github.com/dorajistyle/goyangi/service/likingService/likingRetriever"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/dorajistyle/goyangi/service/userService/userPermission"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/modelHelper"
	"github.com/dorajistyle/goyangi/util/pagination"
	"github.com/dorajistyle/goyangi/util/stringHelper"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// canUserWrite check that user can write a location.
func canUserWrite(c *gin.Context) bool {
	canWrite := false
	_, err := userService.CurrentUser(c)
	if err == nil {
		canWrite = true
	}
	return canWrite
}

// assignRelatedUser assign related user to the location.
func assignRelatedUser(location *model.Location) {
	var tempUser model.User
	if db.ORM.Model(&location).Select(config.UserPublicFields).Related(&tempUser).RecordNotFound() {
		log.Warn("user is not found.")
	}
	location.Author = model.PublicUser{User: &tempUser}
}

// CreateLocation creates a location.
func CreateLocation(c *gin.Context) (model.Location, int, error) {
	var form LocationForm
	c.BindWith(&form, binding.Form)
	log.Debugf("struct map : %s\n", form)
	user, _ := userService.CurrentUser(c)
	form.UserId = user.Id
	location := model.Location{}
	modelHelper.AssignValue(&location, &form)
	if db.ORM.Create(&location).Error != nil {
		return location, http.StatusInternalServerError, errors.New("User is not created.")
	}
	return location, http.StatusCreated, nil
}

// RetrieveLocation retrieves a location.
func RetrieveLocation(c *gin.Context) (model.Location, bool, int64, int, error) {
	var location model.Location
	var currentUserId int64
	isAuthor := false
	id := c.Params.ByName("id")
	if db.ORM.First(&location, "id = ?", id).RecordNotFound() {
		return location, isAuthor, currentUserId, http.StatusNotFound, errors.New("Location is not found.")
	}
	currentUser, err := userService.CurrentUser(c)
	if err == nil {
		currentUserId = currentUser.Id
		isAuthor = currentUser.Id == location.UserId
	}
	assignRelatedUser(&location)
	var commentList model.CommentList
	comments, currentPage, hasPrev, hasNext, _ := commentService.RetrieveComments(location)
	commentList.Comments = comments
	commentService.SetCommentPageMeta(&commentList, currentPage, hasPrev, hasNext, location.CommentCount)
	location.CommentList = commentList
	var likingList model.LikingList
	likings, currentPage, hasPrev, hasNext, _ := likingRetriever.RetrieveLikings(location)
	likingList.Likings = likings
	currentUserlikedCount := db.ORM.Model(&location).Where("id =?", currentUserId).Association("Likings").Count()
	log.Debugf("Current user like count : %d", currentUserlikedCount)
	likingMeta.SetLikingPageMeta(&likingList, currentPage, hasPrev, hasNext, location.LikingCount, currentUserlikedCount)
	location.LikingList = likingList
	return location, isAuthor, currentUserId, http.StatusOK, nil
}

// RetrieveLocations retrieves locations.
func RetrieveLocations(c *gin.Context) ([]model.Location, bool, int, bool, bool, int, error) {
	var locations []model.Location
	var locationCount, locationPerPage int
	filterQuery := c.Request.URL.Query().Get("filter")
	locationPerPage = config.LocationPerPage
	filter := &LocationFilter{}
	whereBuffer := new(bytes.Buffer)
	whereValues := []interface{}{}
	if len(filterQuery) > 0 {
		log.Debugf("retrieve Locations filter : %s\n", filterQuery)
		json.Unmarshal([]byte(filterQuery), &filter)
		if filter.UserId > 0 {
			stringHelper.Concat(whereBuffer, "user_id = ?")
			whereValues = append(whereValues, filter.UserId)
			log.Debugf("userId : %d\n", filter.UserId)
		}
		if filter.LocationPerPage > 0 {
			locationPerPage = filter.LocationPerPage
			log.Debugf("locationPerPage : %d\n", filter.LocationPerPage)
		}
	} else {
		log.Debug("no filters found.\n")
	}
	log.Debugf("filterQuery %v.\n", filterQuery)
	log.Debugf("filter %v.\n", filter)
	whereStr := whereBuffer.String()
	db.ORM.Model(model.Location{}).Where(whereStr, whereValues...).Count(&locationCount)
	offset, currentPage, hasPrev, hasNext := pagination.Paginate(filter.CurrentPage, locationPerPage, locationCount)
	db.ORM.Limit(locationPerPage).Offset(offset).Order(config.LocationOrder).Where(whereStr, whereValues...).Find(&locations)

	return locations, canUserWrite(c), currentPage, hasPrev, hasNext, http.StatusOK, nil
}

// UpdateLocation updates a location.
func UpdateLocation(c *gin.Context) (model.Location, int, error) {
	var location model.Location
	var form LocationForm
	id := c.Params.ByName("id")
	c.BindWith(&form, binding.Form)
	if db.ORM.First(&location, id).RecordNotFound() {
		return location, http.StatusNotFound, errors.New("Location is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, location.UserId)
	if err != nil {
		return location, status, err
	}
	location.Name = form.Name
	location.Address = form.Address
	location.Latitude = form.Latitude
	location.Longitude = form.Longitude
	location.Url = form.Url
	location.Content = form.Content
	if db.ORM.Save(&location).Error != nil {
		return location, http.StatusBadRequest, errors.New("Location is not updated.")
	}
	return location, http.StatusOK, nil
}

// DeleteLocation deletes a location.
func DeleteLocation(c *gin.Context) (int, error) {
	var location model.Location
	id := c.Params.ByName("id")
	if db.ORM.First(&location, id).RecordNotFound() {
		return http.StatusNotFound, errors.New("Location is not found.")
	}
	status, err := userPermission.CurrentUserIdentical(c, location.UserId)
	if err != nil {
		return status, err
	}
	if db.ORM.Delete(&location).Error != nil {
		return http.StatusBadRequest, errors.New("Location is not deleted.")
	}
	return http.StatusOK, nil
}
