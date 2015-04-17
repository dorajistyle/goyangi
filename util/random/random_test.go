package random_test

import (
	. "github.com/dorajistyle/goyangi/util/random"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Random", func() {
	var (
		n            int
		randomString string
	)

	BeforeEach(func() {
		randomString = GenerateRandomString(n)
	})

	Describe("Generate random string", func() {
		Context("when random string generated successfully", func() {
			It("should have length n", func() {
				Expect(len(randomString)).To(Equal(n))
			})

		})
	})
})
