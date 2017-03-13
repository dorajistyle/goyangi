package authHelper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAuthHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authhelper Suite")
}
