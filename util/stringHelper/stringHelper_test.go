package stringHelper_test

import (
	"bytes"
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
		leadingStr      string
		followingStr    string
		concatedString  string
		concatedLength  int
		buffer          *bytes.Buffer
	)

	BeforeEach(func() {
		testStr = "Nowplay is a SDK."
		leadingStr = "Tanguero y "
		followingStr = "Tanguera"
	})

	Describe("Concat string using Concat", func() {
		BeforeEach(func() {
			buffer = new(bytes.Buffer)
			Concat(buffer, testStr)
		})
		Context("when testStr concatenates successfully", func() {
			It("should have length concatedLength", func() {
				Expect(len(buffer.String())).To(Equal(len(testStr)))
			})
		})
	})

	Describe("Concat string using ConcatExist", func() {
		BeforeEach(func() {
			lengthOfStrArr = 1
			strArr = []string{}
			strArr = ConcatExist(strArr, testStr)
			lengthOfTestStr = len(testStr)
			joinedStr = strings.Join(strArr, "")
		})
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

	Describe("Concat string using ConcatString", func() {
		BeforeEach(func() {
			concatedString = ConcatString(leadingStr, followingStr)
			concatedLength = len(concatedString)
		})
		Context("when leadingStr and followingStr concatenates successfully", func() {
			It("should have length concatedLength", func() {
				Expect(concatedLength).To(Equal(len(leadingStr) + len(followingStr)))
			})
		})
	})

})
