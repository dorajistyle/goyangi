package articleService

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/concurrency"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/stringHelper"
	"github.com/dorajistyle/goyangi/util/upload"
)

var (
	Article업로더               = articleUploader()
	ArticleUploader          = articleUploader()
	ArticleCargador          = articleUploader()
	ArticleShàngchuán        = articleUploader()
	Articleзагрузчик         = articleUploader()
	s3UploadPath      string = ""
	user              model.User
)

// articleUploader is a uploader that uploading files and sync articles.
func articleUploader() concurrency.ConcurrencyManager {
	return func(request *http.Request) concurrency.Result {
		atomic.AddInt32(concurrency.BusyWorker, 1)
		var result concurrency.Result
		result.Code = http.StatusOK
		var reader *multipart.Reader
		var err error
		reader, err = request.MultipartReader()
		log.Debug("File upload start.")
		if err != nil {
			log.CheckErrorWithMessage(err, "Uploading failed.")
			result.Code = http.StatusInternalServerError
			result.Error = err
			return result
		}
		var articles []model.Article
		formDataBuffer := make([]byte, 100000)
		fileCount := 0
		sqlStrBuffer := new(bytes.Buffer)
		stringHelper.Concat(sqlStrBuffer, "INSERT INTO article(user_id, title, url, content, image_name, created_at) VALUES ")
		values := []interface{}{}
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			if part.FileName() == "" {
				if part.FormName() != "" {
					log.Debug("formName : " + part.FormName())
					n, err := part.Read(formDataBuffer)
					if err != nil {
						log.Warnf("Multipart read failed. , (Error Detail : %s)", err.Error())
						result.Code = http.StatusBadRequest
						result.Error = err
						return result
					}
					log.Debugf("data : %s ", formDataBuffer)
					err = json.Unmarshal(formDataBuffer[:n], &articles)
					if err != nil {
						log.Warnf("Json unmarshal failed. , (Error Detail : %s)", err.Error())
						result.Code = http.StatusBadRequest
						result.Error = err
						return result
					}
					log.Debugf("err %s", err)
					log.Debugf("article : %v\n", articles)
					log.Debugf("article len : %d\n", len(articles))
				}
				continue
			}
			// err = upload.UploadImageFile("", part)
			err = upload.UploadImageFile(viper.GetString("upload.target"), viper.GetString("app.environment"), s3UploadPath, part)
			if err != nil {
				log.Error("Image uploading failed. : ", err)
				result.Code = http.StatusBadRequest
				result.Error = err
				return result
			}
			if fileCount < len(articles) {
				stringHelper.Concat(sqlStrBuffer, "(?, ?, ?, ?, ?, ?),")
				values = append(values, user.Id, articles[fileCount].Title, articles[fileCount].Url, articles[fileCount].Content, articles[fileCount].ImageName, time.Now())
				// db.ORM.Create(&articles[fileCount])
				fileCount += 1
			}
			log.Debug("File uploaded.")
			log.Infof("File Count : %d\n", fileCount)
		}
		sqlStr := sqlStrBuffer.String()
		sqlStr = sqlStr[0 : len(sqlStr)-1]
		log.Debugf("sqlStr for Article : %s", sqlStr)
		db.ORM.Exec(sqlStr, values...)
		return result
	}
}

// UploadAndSyncArticles uploads images and sync articles.
func UploadAndSyncArticles(c *gin.Context) (int, error) {
	r := c.Request
	s3UploadPath = viper.GetString("upload.path.S3Image") + strconv.FormatInt(int64(user.Id), 10) + "/"
	user, _ = userService.CurrentUser(c)
	concurrency.Concurrent(r, concurrency.ConcurrencyAgent(r, Article업로더, ArticleUploader, ArticleCargador, ArticleShàngchuán, Articleзагрузчик))

	return http.StatusOK, nil
}
