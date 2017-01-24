package redis_test

import (
	"fmt"

	. "github.com/dorajistyle/goyangi/util/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Redis", func() {
	var (
		bestSDK  string
		worstSDK string
	)
	if InitErr != nil {
		fmt.Printf("Redis connection failed : %s\n", InitErr.Error())
		return
	}
	BeforeEach(func() {

		bestSDK = "Nowplay SDK"
		worstSDK = "NowWorst SDK"
		Resource.Append("bestSDKEver", bestSDK)
		Resource.Append("worstSDKEver", worstSDK)
	})
	Describe("get best sdk", func() {

		bestSDKEver, err := Resource.Get("bestSDKEver")
		Context("when redis get a bestSDKEver successfully", func() {
			It("should equals with Nowplay SDK", func() {
				Expect(bestSDKEver).To(Equal(bestSDK))
			})
			It("should have no error", func() {
				Expect(err).To(BeNil())
			})

		})
	})
	Describe("get worst sdk", func() {
		worstSdkEver, err := Resource.Get("worstSDKEver")
		Context("when redis get a worstSDKEver successfully", func() {
			It("should equals with NowWorst SDK", func() {
				Expect(worstSdkEver).To(Equal(worstSDK))
			})
			It("should have no error", func() {
				Expect(err).To(BeNil())
			})

		})
	})
	Describe("get worst sdk after delete key", func() {
		delErr := Resource.Del("worstSDKEver")
		Context("when redis del a worstSDKEver successfully", func() {
			It("should have no error", func() {
				Expect(delErr).To(BeNil())
			})
		})
		worstSDKEver, err := Resource.Get("worstSDKEver")
		Context("when redis get a worstSDKEver successfully", func() {
			It("should equals with NowWorst SDK", func() {
				Expect(worstSDKEver).To(BeEmpty())
			})
			It("should have no error", func() {
				Expect(err).ToNot(BeNil())
			})

		})
	})
	if InitErr == nil {
		Resource.Close()
		Pool.Close()
	}

})
