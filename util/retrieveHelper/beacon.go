package retrieveHelper

import (
	"errors"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/model"
)

func RetriveBeaconWithInfo(key string, types int) (model.Beacon, error) {
	var beacon model.Beacon
	if db.ORM.Where("key = ? and types = ?", key, types).First(&beacon).RecordNotFound() {
		return beacon, errors.New("Beacon does not exist.")
	}
	return beacon, nil
}

// func RetriveBeaconWithInfoLegacy(uuid string, major int, minor int) (model.Beacon, error){
//   var beacon model.Beacon
//   if db.ORM.Where("uuid = ? and major = ? and minor = ?", uuid, major, minor).First(&beacon).RecordNotFound() {
//     return beacon, errors.New("Beacon does not exist.")
//   }
//   return beacon, nil
// }
