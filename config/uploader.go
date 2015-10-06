package config

import "time"

// Constants for uploader.
const (
	UploadLocalPath = "/tmp/upload/"
	UploadS3ImagePath    = "images/"
	//	UploadTarget = LOCAL | S3
	UploadTarget = "LOCAL"
	// UploadTarget  = "S3"
	//	UploadBucket = TEST | PRODUCTION
	UploadBucket  = "TEST"
	
	UploadTimeout = 30 * time.Second
)
