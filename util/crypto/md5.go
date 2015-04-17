package crypto

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func GenerateMD5Hash(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	hash := md5.New()
	hash.Write([]byte(email))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
