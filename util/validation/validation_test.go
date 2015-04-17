package validation_test

import (
	. "github.com/dorajistyle/goyangi/util/validation"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validation", func() {
	var (
		emailValid   string
		emailInvalid string
		isValid      bool
	)

	BeforeEach(func() {
		emailValid = "test@goyangi.github.io"
		emailInvalid = "test#goyangi.github.io"
	})

	Describe("Check that the email address is valid or not", func() {

		Context("when the email checked successfully", func() {
			BeforeEach(func() {
				isValid = EmailValidation(emailValid)
			})

			It("should be valid.", func() {
				Expect(isValid).To(Equal(true))
			})
		})

		Context("when the email checked successfully", func() {
			BeforeEach(func() {
				isValid = EmailValidation(emailInvalid)
			})

			It("should be invalid.", func() {
				Expect(isValid).To(Equal(false))
			})
		})
	})

})
