package locationService

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/form"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/service/likingService"
	"github.com/dorajistyle/goyangi/service/likingService/likingRetriever"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// UpdateLocationLikingCount updates a liking count on article.
func UpdateLocationLikingCount(location *model.Location) (int, error) {
	location.LikingCount = db.ORM.Model(location).Association("Likings").Count()
	if db.ORM.Save(location).Error != nil {
		return http.StatusInternalServerError, errors.New("Location liking's count is not updated.")
	}
	return http.StatusOK, nil
}

// CreateLikingOnLocation creates a liking on location.
func CreateLikingOnLocation(c *gin.Context) (int, error) {
	location := &model.Location{}
	status, err := likingService.CreateLiking(c, location)
	if err != nil {
		return status, err
	}
	status, err = UpdateLocationLikingCount(location)
	if err != nil {
		return status, err
	}
	return http.StatusOK, nil
}

// RetrieveLikingsOnLocations retrieves likings on location.
func RetrieveLikingsOnLocations(c *gin.Context) ([]*model.PublicUser, int, bool, bool, int, int, error) {
	var location model.Location
	var likings []*model.PublicUser
	var retrieveListForm form.RetrieveListForm
	var hasPrev, hasNext bool
	var currentPage, count int
	locationId := c.Params.ByName("id")
	log.Debugf("Liking params : %v", c.Params)
	c.BindWith(&retrieveListForm, binding.Form)
	log.Debugf("retrieveListForm %+v\n", retrieveListForm)
	if db.ORM.First(&location, locationId).RecordNotFound() {
		return likings, currentPage, hasPrev, hasNext, count, http.StatusNotFound, errors.New("Location is not found.")
	}
	likings, currentPage, hasPrev, hasNext, count = likingRetriever.RetrieveLikings(location, retrieveListForm.CurrentPage)
	return likings, currentPage, hasPrev, hasNext, count, http.StatusOK, nil
}

// DeleteLikingOnLocation deletes liking on location.
func DeleteLikingOnLocation(c *gin.Context) (int, error) {
	location := &model.Location{}
	status, err := likingService.DeleteLiking(c, location)
	if err != nil {
		return status, err
	}
	status, err = UpdateLocationLikingCount(location)
	if err != nil {
		return status, err
	}
	return http.StatusOK, nil
}
