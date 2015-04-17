package timeHelper_test

import (
	"time"

	. "github.com/dorajistyle/goyangi/util/timeHelper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TimeHelper", func() {
	var (
		duration  time.Duration
		now       time.Time
		later     time.Time
		isExpired bool
	)

	BeforeEach(func() {
		duration = time.Duration(6) * time.Hour
		now = time.Now()
	})

	Describe("Get the time that few duration later", func() {

		Context("when getting the time successfully", func() {
			BeforeEach(func() {
				later = FewDurationLater(duration)
			})

			It("should be 6 hours later.", func() {
				Expect(later.Hour()).To(Equal(now.Add(duration).Hour()))
			})

		})
	})

	Describe("Check the time is expired", func() {

		Context("when time expiration checked successfully", func() {
			BeforeEach(func() {
				isExpired = IsExpired(now)
			})

			It("should be expired.", func() {
				Expect(isExpired).To(Equal(true))
			})

		})
	})

})
