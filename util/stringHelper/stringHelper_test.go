package stringHelper_test

import (
	"strings"

	. "github.com/dorajistyle/goyangi/util/stringHelper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringHelper", func() {
	var (
		testStr         string
		strArr          []string
		lengthOfTestStr int
		lengthOfStrArr  int
		joinedStr       string
	)

	BeforeEach(func() {
		testStr = "Goyangi means cat."
		strArr = []string{}
		lengthOfTestStr = len(testStr)
		lengthOfStrArr = 1

		strArr = ConcatExist(strArr, testStr)
		joinedStr = strings.Join(strArr, "")
	})

	Describe("Concat string using ConcatExist", func() {
		Context("when testStr appended to strArr successfully", func() {
			It("should have length lengthOfStrArr", func() {
				Expect(len(strArr)).To(Equal(lengthOfStrArr))
			})
		})
		Context("when strArr joined successfully", func() {
			It("should have length lengthOfTestStr", func() {
				Expect(len(joinedStr)).To(Equal(lengthOfTestStr))
			})
		})
	})
})
