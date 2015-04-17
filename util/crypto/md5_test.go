package crypto_test

import (
	. "github.com/dorajistyle/goyangi/util/crypto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Md5", func() {
	var (
		emailOne string
		emailTwo string
	)

	BeforeEach(func() {
		emailOne = "testOne@test.com"
		emailTwo = "testTwo@test.com"
	})

	Describe("Generate MD5 hash from email", func() {
		Context("when the MD5 hash of emailOne generated successfully", func() {
			It("should be a correct hash", func() {
				Expect(GenerateMD5Hash(emailOne)).To(Equal("d45cfb8dfe120832336109537e52d1c7"))
			})
		})
		Context("when the MD5 hash of emailTwo generated successfully", func() {
			It("should be a correct hash", func() {
				Expect(GenerateMD5Hash(emailTwo)).To(Equal("327e27ad95910153a96d255833c45328"))
			})
		})
	})
})
