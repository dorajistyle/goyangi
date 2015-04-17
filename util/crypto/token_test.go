package crypto_test

import (
	. "github.com/dorajistyle/goyangi/util/crypto"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Token", func() {
	var (
		num32  int
		num64  int
		num128 int
		token  string
		err    error
	)

	BeforeEach(func() {
		num32 = 32
		num64 = 64
		num128 = 128
	})

	Describe("Generate random token", func() {
		Context("when random token generated successfully", func() {
			BeforeEach(func() {
				token, err = GenerateRandomToken16()
			})
			It("should have length 32", func() {
				Expect(len(token)).To(Equal(num32))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when random token generated successfully", func() {
			BeforeEach(func() {
				token, err = GenerateRandomToken32()
			})
			It("should have length 64", func() {
				Expect(len(token)).To(Equal(num64))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when random token generated successfully", func() {
			BeforeEach(func() {
				token, err = GenerateRandomToken(num64)
			})
			It("should have length 128", func() {
				Expect(len(token)).To(Equal(num128))
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})
})
