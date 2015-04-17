package pagination_test

import (
	. "github.com/dorajistyle/goyangi/util/pagination"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pagination", func() {
	var (
		offset      int
		currentPage int
		perPage     int
		total       int
		hasPrev     bool
		hasNext     bool
	)

	BeforeEach(func() {
		currentPage = 2
		perPage = 10
		total = 17
	})

	Describe("Paginate items", func() {
		Context("when ", func() {
			BeforeEach(func() {
				offset, currentPage, hasPrev, hasNext = Paginate(currentPage, perPage, total)
			})

			It("should be 5", func() {
				Expect(offset).To(Equal(10))
			})

			It("should be 2", func() {
				Expect(currentPage).To(Equal(2))
			})

			It("should be true", func() {
				Expect(hasPrev).To(Equal(true))
			})

			It("should be false", func() {
				Expect(hasNext).To(Equal(false))
			})
		})
	})

})
