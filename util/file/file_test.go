package file_test

import (
	"bytes"

	. "github.com/dorajistyle/goyangi/util/file"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
	var (
		filename string
		err      error
		wb       *bytes.Buffer
	)

	BeforeEach(func() {
		filename = "testfile.txt"
		wb = new(bytes.Buffer)
	})

	Describe("Create a directory", func() {
		Context("when the uploadPath is created successfully", func() {
			BeforeEach(func() {
				_, err = UploadPath()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
	Describe("Save a file to local directory", func() {
		Context("when the file saved successfully", func() {
			BeforeEach(func() {
				err = SaveLocal(filename, wb)
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
	Describe("Delete a file to local directory", func() {
		Context("when the file deleted successfully", func() {
			BeforeEach(func() {
				err = SaveLocal(filename, wb)
				err = DeleteLocal(filename)
			})
			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
