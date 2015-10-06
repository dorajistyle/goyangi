package uploadService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/dorajistyle/goyangi/db"
	"github.com/dorajistyle/goyangi/service/userService"
	"github.com/gin-gonic/gin"

	"github.com/dorajistyle/goyangi/config"
	"github.com/dorajistyle/goyangi/model"
	"github.com/dorajistyle/goyangi/util/aws"
	"github.com/dorajistyle/goyangi/util/file"
	"github.com/dorajistyle/goyangi/util/image"
	"github.com/dorajistyle/goyangi/util/log"
	"github.com/dorajistyle/goyangi/util/stringHelper"
)

type UploadStatus bool
type Uploader func(reader *multipart.Reader) UploadStatus

// Kinds of uploader.
const (
	KindUploaderBasic   = 1
	KindUploaderArticle = 2
)

var (
	user            model.User
	s3UploadPath    string
	uploaded        *uint32 = new(uint32)
	workingUploader *int32  = new(int32)

	업로더           = basicUploader()
	BasicUploader = basicUploader()
	Cargador      = basicUploader()
	Shàngchuán    = basicUploader()
	загрузчик     = basicUploader()

	Article업로더        = articleUploader()
	ArticleUploader   = articleUploader()
	ArticleCargador   = articleUploader()
	ArticleShàngchuán = articleUploader()
	Articleзагрузчик  = articleUploader()
)

// Upload uploads file to a storage.
func Upload(reader *multipart.Reader, kind int) {
	c := make(chan UploadStatus)
	switch kind {
	case KindUploaderBasic:
		go func() {
			c <- UploadAgent(reader, 업로더, BasicUploader, Cargador, Shàngchuán, загрузчик)
		}()
	case KindUploaderArticle:
		go func() {
			c <- UploadAgent(reader, Article업로더, ArticleUploader, ArticleCargador, ArticleShàngchuán, Articleзагрузчик)
		}()
	default:
		go func() {
			c <- UploadAgent(reader, 업로더, BasicUploader, Cargador, Shàngchuán, загрузчик)
		}()
	}

	timeout := time.After(config.UploadTimeout)
	select {
	case <-c:
		workingNow := atomic.AddInt32(workingUploader, -1)
		log.Debugf("All files are uploaded. Working uploader count : %d", workingNow)
		return
	case <-timeout:
		fmt.Println("timed out")
		return
	}
}

// UploadAgent is loadbalancer of uploader.
func UploadAgent(reader *multipart.Reader, replicas ...Uploader) UploadStatus {
	for {
		workingNow := atomic.LoadInt32(workingUploader)
		if len(replicas) > int(workingNow) {
			break
		}
		time.Sleep(time.Second)
		// log.Debugf("working uploader count full (workingUploader/replicas)  : (%d/%d) ", workingNow, len(replicas))
	}
	c := make(chan UploadStatus)

	uploaderReplica := func(i int) {
		c <- replicas[i](reader)
	}
	workingNow := atomic.LoadInt32(workingUploader)
	log.Debugf("workingNow, len(replicas) : %d %d", workingNow, len(replicas))

	go uploaderReplica(int(workingNow))
	// go uploaderReplica(0)
	return <-c
}

// articleUploader is a uploader that uploading files and sync articles.
func articleUploader() Uploader {
	return func(reader *multipart.Reader) UploadStatus {
		atomic.AddInt32(workingUploader, 1)
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
					log.Debugf("n, err %d %s", n, err)
					log.Debugf("data : %s ", formDataBuffer)
					err = json.Unmarshal(formDataBuffer[:n], &articles)
					log.Debugf("err %s", err)
					log.Debugf("article : %v\n", articles)
					log.Debugf("article len : %d\n", len(articles))
				}
				continue
			}
			UploadImageFile(part)
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
		return UploadStatus(true)
	}
}

// basicUploader is a uploader that uploading files.
func basicUploader() Uploader {
	return func(reader *multipart.Reader) UploadStatus {
		atomic.AddInt32(workingUploader, 1)
		for {
			part, err := reader.NextPart()

			uploadedNow := atomic.AddUint32(uploaded, 1)
			log.Debugf("count %d", uploadedNow)
			if err == io.EOF {
				log.Warn("End of file.")
				break
			}
			if part.FileName() == "" {
				log.Warn("File name is empty.")
				continue
			}
			UploadImageFile(part)
			log.Debug("File uploaded.")
		}
		log.Debug("Iteration done.")
		return UploadStatus(true)
	}
}

// UploadImageFile uploads an image file to a storage.
func UploadImageFile(part *multipart.Part) {

	mediatype, _, _ := mime.ParseMediaType(part.Header.Get("Content-Type"))
	log.Debugf("params %s", mediatype)
	log.Debug("fileName : " + part.FileName())
	inbuf, err := ioutil.ReadAll(part)
	if err != nil {
		log.CheckErrorWithMessage(err, "Image read failed.")
	}
	// Image resize is a bottleneck. How can we improve this?
	// https://github.com/fawick/speedtest-resize said vipsthumbnail is fastest one.
	// Currenctly goyangi uses vips(https://github.com/DAddYE/vips).

	// dst, _ := image.ResizeMedium(mediatype, bytes.NewReader(inBuf))
	var dst *bytes.Buffer
	buf, err := image.ResizeMediumVips(inbuf)
	if err != nil {
		log.CheckErrorWithMessage(err, "Image resizing failed.")
		dst = bytes.NewBuffer(inbuf)
	} else {
		dst = bytes.NewBuffer(buf)
	}

	// var thumbDst *bytes.Buffer
	// thumbBuf, err := image.ResizeThumbnailVips(buf)
	// if err != nil {
	// 	log.CheckErrorWithMessage(err, "Image thumbnailing failed.")
	// 	thumbDst = bytes.NewBuffer(buf)
	// } else {
	// 	thumbDst = bytes.NewBuffer(thumbBuf)
	// }

	switch config.UploadTarget {
	case "LOCAL":
		err = file.SaveLocal(part.FileName(), dst)
		// err = file.SaveLocal(part.FileName()+"Thumbnail", thumbDst)
	case "S3":
		err = aws.PutToMyPublicBucket(s3UploadPath, part.FileName(), dst, mediatype)
		// err = aws.PutToMyPublicBucket("images/", part.FileName()+"Thumbnail", thumbDst, mediatype)
	}
	if err != nil {
		log.CheckErrorWithMessage(err, "Uploading failed.")
	}
}

// UploadImages uploads images to a storage.
func UploadImages(c *gin.Context) (int, error) {
	r := c.Request
	reader, err := r.MultipartReader()
	user, _ = userService.CurrentUser(c)
	s3UploadPath = config.UploadS3Path + strconv.FormatInt(user.Id, 10) + "/"
	if err != nil {
		return http.StatusInternalServerError, err
	}
	Upload(reader, KindUploaderBasic)
	return http.StatusOK, nil
}

// UploadAndSyncArticles uploads images and sync articles.
func UploadAndSyncArticles(c *gin.Context) (int, error) {
	r := c.Request
	reader, err := r.MultipartReader()
	user, _ = userService.CurrentUser(c)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	Upload(reader, KindUploaderArticle)

	return http.StatusOK, nil
}
