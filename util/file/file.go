package file

import (
	"bytes"
	"io"
	"os"

	"github.com/dorajistyle/goyangi/util/config"
)

// UploadPath gets upload path.
func UploadPath() (string, error) {
	uploadPath := config.UploadLocalPath
	err := os.MkdirAll(uploadPath, 0777)
	return uploadPath, err
}

// SaveLocal save a file locally.
func SaveLocal(fileName string, wb *bytes.Buffer) error {
	uploadPath, err := UploadPath()
	if err != nil {
		return err
	}
	var out *os.File
	out, err = os.Create(uploadPath + fileName)
	defer out.Close()
	if err != nil {
		return err
	}
	_, err = io.Copy(out, wb)
	return err
}

// DeleteLocal delete a file locally.
func DeleteLocal(fileName string) error {
	uploadPath, err := UploadPath()
	if err != nil {
		return err
	}
	err = os.Remove(uploadPath + fileName)
	return err
}
