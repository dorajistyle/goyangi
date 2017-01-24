package redis

import (
	"errors"
	"net/http"

	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/stringHelper"
	"github.com/gin-gonic/gin"
)

// responseGenerator types take an *gin.Context and return a string, an int and an error value.
type responseGenerator func(c *gin.Context) (string, int, error)

// CacheResponse caches a response if redis server is available.
func CacheResponse(c *gin.Context, keyPrefix string, keyBody string, resGenerator responseGenerator) (string, int, error) {
	var encrypted string
	var status int
	var resErr error
	var cacheStr string
	var cacheKey string
	var cacheErr error
	cacheErr = errors.New("It is a default error of cache")
	cacheKey = stringHelper.ConcatString(keyPrefix, keyBody)
	if InitErr == nil {
		cacheStr, cacheErr = Resource.Get(cacheKey)
	}
	if cacheErr != nil {
		if resGenerator == nil {
			log.Debug("resGenerator is nil")
			return encrypted, status, cacheErr
		}
		encrypted, status, resErr = resGenerator(c)
		if InitErr == nil {
			log.Debugf("cacheKey : %s\n", cacheKey)
			appendErr := Resource.Append(cacheKey, encrypted)
			if appendErr != nil {
				log.Error(appendErr.Error())
			}
		}
	} else {
		log.Debugf("It has cached response. Key : %s\n", cacheKey)
		encrypted = cacheStr
		status = http.StatusOK
	}
	return encrypted, status, resErr
}
