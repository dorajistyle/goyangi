package redis_test

import (
	. "github.com/dorajistyle/goyangi/util/redis"
	viper "github.com/dorajistyle/goyangi/util/viper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	viper.LoadConfig()
}

var _ = Describe("Redis", func() {
	var (
		bestSDK  string
		worstSDK string
	)

	BeforeEach(func() {

		bestSDK = "Goyangi SDK"
		worstSDK = "Goyak SDK"
		Append("bestSDKEver", bestSDK)
		Append("worstSDKEver", worstSDK)
	})
	Describe("get best sdk", func() {

		bestSDKEver, err := Get("bestSDKEver")
		// fmt.Println("bestSDKEver: %s", bestSDKEver)
		Context("when redis get a bestSDKEver successfully", func() {
			It("should equals with Goyangi SDK", func() {
				Expect(bestSDKEver).To(Equal(bestSDK))
			})
			It("should have no error", func() {
				Expect(err).To(BeNil())
			})

		})
	})
	Describe("get worst sdk", func() {
		worstSdkEver, err := Get("worstSDKEver")
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
		delErr := Del("worstSDKEver")
		Context("when redis del a worstSDKEver successfully", func() {
			It("should have no error", func() {
				Expect(delErr).To(BeNil())
			})
		})
		worstSDKEver, err := Get("worstSDKEver")
		Context("when redis get a worstSDKEver successfully", func() {
			It("should equals with NowWorst SDK", func() {
				Expect(worstSDKEver).To(BeEmpty())
			})
			It("should have no error", func() {
				Expect(err).ToNot(BeNil())
			})

		})
	})
})
