package retrieveHelper

import (
	"errors"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
)

func RetriveAppWithAppKey(appKey string) (model.App, error) {
	var app model.App
	if db.ORM.Where("key=?", appKey).First(&app).RecordNotFound() {
		return app, errors.New("App does not exist.")
	}
	return app, nil
}
