package upload

import (
  "bytes"
  "io/ioutil"
  "mime"
  "mime/multipart"
  "strings"
  "path/filepath"
  "github.com/dorajistyle/goyangi/config"
  "github.com/dorajistyle/goyangi/util/aws"
  "github.com/dorajistyle/goyangi/util/file"
  "github.com/dorajistyle/goyangi/util/image"
  "github.com/dorajistyle/goyangi/util/log"
  )

  // UploadImageFile uploads an image file to a storage.
  func UploadImageFile(s3UploadPath string, part *multipart.Part) error {
    mediatype, _, _ := mime.ParseMediaType(part.Header.Get("Content-Type"))
    log.Debugf("params %s", mediatype)
    log.Debug("fileName : " + part.FileName())
    inbuf, err := ioutil.ReadAll(part)
    if err != nil {
      // log.CheckErrorWithMessage(err, "Image read failed.")
      return err
    }
    // Image resize is a bottleneck. How can we improve this?
    // https://github.com/fawick/speedtest-resize said vipsthumbnail is fastest one.
    // Currenctly goyangi uses vips(https://github.com/DAddYE/vips).
    // dst, _ := image.ResizeMedium(mediatype, bytes.NewReader(inBuf))
    var dst, dstLarge, dstMedium, dstThumbnail *bytes.Buffer
      dst = bytes.NewBuffer(inbuf)
      buf, err := image.ResizeLargeVips(inbuf)
      if err != nil {
        // log.CheckErrorWithMessage(err, "Image resizing failed.")
        log.Errorf("Image large resizing failed. %s", err.Error())
        dstLarge = nil
        } else {
          dstLarge = bytes.NewBuffer(buf)
          mbuf, err := image.ResizeMediumVips(buf)
          if err != nil {
            dstMedium = nil
            log.Errorf("Image medium resizing failed. %s", err.Error())
          } else {
            dstMedium = bytes.NewBuffer(mbuf)
            tbuf, err := image.ResizeThumbnailVips(mbuf)
            if err != nil {
              dstThumbnail = nil
              log.Errorf("Image small resizing failed. %s", err.Error())
            } else {
              dstThumbnail = bytes.NewBuffer(tbuf)
            }
          }
        }

      // var thumbDst *bytes.Buffer
      // thumbBuf, err := image.ResizeThumbnailVips(buf)
      // if err != nil {
      // 	log.CheckErrorWithMessage(err, "Image thumbnailing failed.")
      // 	thumbDst = bytes.NewBuffer(buf)
      // } else {
      // 	thumbDst = bytes.NewBuffer(thumbBuf)
      // }
      basename := part.FileName()
      ext := filepath.Ext(basename)
      name := strings.TrimSuffix(basename, ext)
      originName := basename
      largeName := name+"_large"+ext
      mediumName := name+"_medium"+ext
      thumbnailName := name+"_thumbnail"+ext
      switch config.UploadTarget {
        case "LOCAL":
        err = file.SaveLocal(originName, dst)
        if dstLarge != nil {
          err = file.SaveLocal(largeName, dstLarge)
        }
        if dstMedium != nil {
          err = file.SaveLocal(mediumName, dstMedium)
        }
        if dstThumbnail != nil {
          err = file.SaveLocal(thumbnailName, dstThumbnail)
        }

        case "S3":
          switch config.Environment {
            case "DEVELOPMENT":
              fallthrough
            case "TEST":
                err = aws.PutToMyPublicTestBucket(s3UploadPath, originName, dst, mediatype)
                if dstLarge != nil {
                  err = aws.PutToMyPublicTestBucket(s3UploadPath, largeName, dstLarge, mediatype)
                }
                if dstMedium != nil {
                  err = aws.PutToMyPublicTestBucket(s3UploadPath, mediumName, dstMedium, mediatype)
                }
                if dstThumbnail != nil {
                  err = aws.PutToMyPublicTestBucket(s3UploadPath, thumbnailName, dstThumbnail, mediatype)
                }
            case "PRODUCTION":
            err = aws.PutToMyPublicBucket(s3UploadPath, originName, dst, mediatype)
            if dstLarge != nil {
              err = aws.PutToMyPublicBucket(s3UploadPath, largeName, dstLarge, mediatype)
            }
            if dstMedium != nil {
              err = aws.PutToMyPublicBucket(s3UploadPath, mediumName, dstMedium, mediatype)
            }
            if dstThumbnail != nil {
              err = aws.PutToMyPublicBucket(s3UploadPath, thumbnailName, dstThumbnail, mediatype)
            }
          }

        // err = aws.PutToMyPublicBucket("images/", part.FileName()+"Thumbnail", thumbDst, mediatype)
      }
      if err != nil {
        // log.CheckErrorWithMessage(err, "Uploading failed.")
        return err
      }
      return nil
    }
