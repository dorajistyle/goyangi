package ginHelper

import (
    "github.com/gin-gonic/gin"
)

 // PostFormArray returns a slice of strings for a given form key, plus
 // a boolean value whether at least one value exists for the given key.
 func PostFormArray(c *gin.Context, key string) ([]string, bool) {
  	req := c.Request
 	req.ParseForm()
  	req.ParseMultipartForm(32 << 20) // 32 MB
  	if values := req.PostForm[key]; len(values) > 0 {
 		return values, true
  	}
  	if req.MultipartForm != nil && req.MultipartForm.File != nil {
  		if values := req.MultipartForm.Value[key]; len(values) > 0 {
 			return values, true
  		}
  	}
 	return []string{}, false
  }