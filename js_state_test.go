package jsgo_test

import (
	. "github.com/go-zero/jsgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsState", func() {

	Context("with a new one", func() {
		var state *JsState

		BeforeEach(func() {
			state = NewJsState()
		})

		AfterEach(func() {
			state.Free()
		})

		It("Shouldn't be null", func() {
			Expect(state).ToNot(BeNil())
		})

		It("Should be able to run a piece of code", func() {
			_, err := state.DoString("1")
			Expect(err).ToNot(HaveOccurred())
		})

		It("Should return an error when try to run an invalid code", func() {
			_, err := state.DoString("a")
			Expect(err).To(HaveOccurred())

			jsError, _ := err.(*JsError)
			Expect(jsError.Name).To(Equal("ReferenceError"))
		})

	})

})
