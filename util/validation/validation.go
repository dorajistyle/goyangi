package validation

import (
	"regexp"

	"github.com/dorajistyle/goyangi/util/log"
)

const EMAIL_REGEX = `(\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3})`

func EmailValidation(email string) bool {
	exp, err := regexp.Compile(EMAIL_REGEX)
	if regexpCompiled := log.CheckError(err); regexpCompiled {
		if exp.MatchString(email) {
			return true
		}
	}
	return false
}
