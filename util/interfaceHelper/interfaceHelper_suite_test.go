package interfaceHelper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInterfaceHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "InterfaceHelper Suite")
}
