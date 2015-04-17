package aws

import (
	"bytes"

	"github.com/dorajistyle/goyangi/config"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
)

// Connection create connection of S3.
func Connection(auth aws.Auth, region aws.Region) *s3.S3 {
	return s3.New(auth, region)
}

// Bucket get bucket of S3 via bucket name.
func Bucket(connection *s3.S3, bucketName string) *s3.Bucket {
	return connection.Bucket(bucketName)
}

// MyBucket get bucket from config.
func MyBucket() *s3.Bucket {
	return Bucket(Connection(Auth(), Region(config.AWSS3RegionName)), config.AWSS3BucketName)
}

// List get list from bucket.
func List(bucket *s3.Bucket, prefix, delim, marker string, max int) (*s3.ListResp, error) {
	return bucket.List(prefix, delim, marker, max)
}

// MyBucketList get list from MyBucket.
func MyBucketList(prefix, delim, marker string, max int) (*s3.ListResp, error) {
	return List(MyBucket(), prefix, delim, marker, max)
}

// PutToMyBucket put a file to a bucket.
func PutToMyBucket(prefix string, keyname string, wb *bytes.Buffer, contentType string, aclType string) error {
	acl := s3.ACL(aclType)
	return MyBucket().Put(prefix+keyname, wb.Bytes(), contentType, acl, s3.Options{})
}

// PutToMyPublicBucket put a file to the MyBucket.
func PutToMyPublicBucket(subdir string, keyname string, wb *bytes.Buffer, contentType string) error {
	return PutToMyBucket(config.AWSS3BucketPrefix, subdir+keyname, wb, contentType, "public-read")
}

// DelFromMyBucket delete a file from a bucket.
func DelFromMyBucket(prefix string, keyname string) error {
	return MyBucket().Del(config.AWSS3BucketPrefix + prefix + keyname)
}
