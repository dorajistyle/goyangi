package config

import "time"

// Constants for uploader.
const (
	UploadLocalPath = "/tmp/upload/"
	UploadS3Path    = "images/"
	//	UploadTarget = LOCAL | S3
	UploadTarget = "LOCAL"
	// UploadTarget  = "S3"
	UploadTimeout = 30 * time.Second
)
