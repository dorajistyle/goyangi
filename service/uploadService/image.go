package uploadService

import (
  "io"
  "mime/multipart"
  "net/http"
  "strconv"
  "sync/atomic"

  "github.com/dorajistyle/goyangi/service/userService"
  "github.com/gin-gonic/gin"

  "github.com/dorajistyle/goyangi/config"
  "github.com/dorajistyle/goyangi/model"
  "github.com/dorajistyle/goyangi/util/log"
  "github.com/dorajistyle/goyangi/util/upload"
  "github.com/dorajistyle/goyangi/util/concurrency"
  )


  var (
    업로더           = imageUploader()
    ImageUploader = imageUploader()
    Cargador      = imageUploader()
    Shàngchuán    = imageUploader()
    загрузчик     = imageUploader()
    s3UploadPath string = ""
    user model.User
    )



// imageUploader is a uploader that uploading files.
func imageUploader() concurrency.ConcurrencyManager {
	return func(reader *multipart.Reader) concurrency.Result {
		atomic.AddInt32(concurrency.BusyWorker, 1)
    var result concurrency.Result
    result.Code = http.StatusOK
		for {
			part, err := reader.NextPart()

			uploadedNow := atomic.AddUint32(concurrency.Done, 1)
			log.Debugf("count %d", uploadedNow)
			if err == io.EOF {
				log.Debug("End of file.")
				break
			}
      if part.FileName() == "" {
        log.Debug("File name is empty.")
        continue
      }
			err = upload.UploadImageFile(s3UploadPath, part)
      if err != nil {
        log.Error("Image uploading failed. : "+err.Error())
        result.Code = http.StatusBadRequest
        result.Error = err
        return result
      }
			log.Debug("File uploaded.")
		}
		log.Debug("Iteration concurrency.Done.")
		return result
	}
}

// UploadImages uploads images to a storage.
func UploadImages(c *gin.Context) (int, error) {
	r := c.Request
	reader, err := r.MultipartReader()
	user, _ = userService.CurrentUser(c)
	s3UploadPath = config.UploadS3ImagePath + strconv.FormatInt(user.Id, 10) + "/"
	if err != nil {
		return http.StatusInternalServerError, err
	}
	concurrency.Concurrent(reader, concurrency.ConcurrencyAgent(reader, 업로더, ImageUploader, Cargador, Shàngchuán, загрузчик))
	return http.StatusOK, nil
}
