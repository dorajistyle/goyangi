package image_test

import (
	"bytes"
	"image"
	"io"
	"os"

	"github.com/disintegration/gift"
	. "github.com/dorajistyle/goyangi/util/image"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func ReadFile(filename string) (io.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		println("File read failed.")
		return file, err
	}
	// defer f.Close()
	println("File read success.")

	return file, err
}

var _ = Describe("Image", func() {
	var (
		thumbnailFilter *gift.GIFT
		mediumFilter    *gift.GIFT
		err             error
		wb              *bytes.Buffer
		srcImg          *image.Gray
		file            *os.File
		jpegFilename    string
		pngFilename     string
		gifFilename     string
	)

	BeforeEach(func() {
		thumbnailFilter = ThumbnailFilter()
		mediumFilter = MediumFilter()
		srcImg = image.NewGray(image.Rect(0, 0, 10, 20))
		wb = new(bytes.Buffer)
		jpegFilename = "testdata/cat.jpg"
		pngFilename = "testdata/cat.png"
		gifFilename = "testdata/cat.gif"
	})

	Describe("Apply filter", func() {
		Context("when filter applied successfully", func() {
			It("should it's bound be", func() {
				Expect(ApplyFilter(srcImg, thumbnailFilter).Bounds()).To(Equal(thumbnailFilter.Bounds(srcImg.Bounds())))
			})
		})
	})
	Describe("Handle jpeg", func() {
		BeforeEach(func() {
			file, err = os.Open(jpegFilename)
		})
		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when jpeg parsed with filter successfully", func() {

			BeforeEach(func() {
				err = ParseJpeg(wb, file, mediumFilter)
				defer file.Close()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when jpeg resized with filter successfully", func() {

			BeforeEach(func() {
				err = ResizeJpeg(wb, file, mediumFilter, 50, 50)
				defer file.Close()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

	Describe("Handle png", func() {
		BeforeEach(func() {
			file, err = os.Open(pngFilename)
		})
		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when png parsed with filter successfully", func() {

			BeforeEach(func() {
				err = ParsePng(wb, file, mediumFilter)
				defer file.Close()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when png resized with filter successfully", func() {

			BeforeEach(func() {
				err = ResizePng(wb, file, mediumFilter, 50, 50)
				defer file.Close()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

	Describe("Handle gif", func() {
		BeforeEach(func() {
			file, err = os.Open(gifFilename)
		})
		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when gif parsed with filter successfully", func() {

			BeforeEach(func() {
				err = ParseGif(wb, file, mediumFilter)
				defer file.Close()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("when gif resized with filter successfully", func() {

			BeforeEach(func() {
				err = ResizeGif(wb, file, mediumFilter, 50, 50)
				defer file.Close()
			})

			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})
		})

	})

})
