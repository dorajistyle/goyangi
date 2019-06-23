package modelHelper_test

import (
	. "github.com/dorajistyle/goyangi/util/modelHelper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestModel struct {
	ID    int64
	Title string
	Order int
}
type TestForm struct {
	ID    int64
	Title string
	Level int
}

var _ = Describe("ModelHelper", func() {
	var (
		model *TestModel
		form  *TestForm
	)

	BeforeEach(func() {
		model = &TestModel{}
		form = &TestForm{ID: 4, Title: "goyangi, the beast."}
	})

	Describe("Check the interface", func() {
		Context("when the type is zero", func() {
			It("should equal to false", func() {
				Expect(IsZeroOfUnderlyingType(model)).To(Equal(false))
			})
		})
	})
	Describe("Assign values of form to model", func() {
		Context("when values assigned successfully", func() {
			BeforeEach(func() {
				AssignValue(model, form)
			})
			It("should have same ID", func() {
				Expect(model.ID).To(Equal(form.ID))
			})

			It("should have same title", func() {
				Expect(model.Title).To(Equal(form.Title))
			})
		})
	})
})
