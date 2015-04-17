package stringHelper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStringHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringHelper Suite")
}
