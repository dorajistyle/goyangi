package redis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	"testing"
)

func TestRedis(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redis Suite")
}
