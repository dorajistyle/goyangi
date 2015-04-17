package timeHelper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTimeHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TimeHelper Suite")
}
