package interfaceHelper_test

import (
	. "github.com/dorajistyle/goyangi/util/interfaceHelper"
	viper "github.com/dorajistyle/goyangi/util/viper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	viper.LoadConfig()
}

var _ = Describe("InterfaceHelper", func() {
	var (
		// inputInt64 interface{}
		valueFloat32 float32
		valueFloat64 float64
		valueInt     int
		valueInt8    int8
		valueInt16   int16
		valueInt32   int32
		valueInt64   int64
		valueUInt    uint
		valueUInt8   uint8
		valueUInt16  uint16
		valueUInt32  uint32
		valueUInt64  uint64
		valueString  string
		outputInt64  int64
	)

	BeforeEach(func() {
		valueFloat32 = 124.00
		valueFloat64 = 124.00
		valueInt = 124
		valueInt8 = 124
		valueInt16 = 124
		valueInt32 = 124
		valueInt64 = 124
		valueUInt = 124
		valueUInt8 = 124
		valueUInt16 = 124
		valueUInt32 = 124
		valueUInt64 = 124
		valueString = "124"
	})

	Describe("Get int64 value from interface{}", func() {

		Context("when getting the int64 value successfully", func() {
			BeforeEach(func() {

			})
			It("that casted from float32 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueFloat32)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from float64 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueFloat64)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from int type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueInt)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from int8 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueInt8)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from int16 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueInt16)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from int32 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueInt32)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from int64 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueInt64)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from uint type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueUInt)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from uint8 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueUInt8)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from uint16 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueUInt16)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from uint32 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueUInt32)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from uint64 type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueUInt64)
				Expect(outputInt64).To(Equal(valueInt64))
			})
			It("that casted from string type should be valueInt64.", func() {
				outputInt64, _ = GetInt64(valueString)
				Expect(outputInt64).To(Equal(valueInt64))
			})

		})
	})
})
