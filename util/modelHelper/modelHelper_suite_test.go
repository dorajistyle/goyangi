package modelHelper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestModelHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ModelHelper Suite")
}
