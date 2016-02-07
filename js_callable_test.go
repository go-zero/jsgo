package jsgo_test

import (
	. "github.com/go-zero/jsgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsCallable", func() {
	state := NewJsState()

	It("Should be callable", func() {
		value, err := state.DoString("new Function('return 42')")
		Expect(err).ToNot(HaveOccurred())

		callable, ok := value.(*JsCallable)
		Expect(ok).To(BeTrue())

		result, err := callable.Call()
		Expect(err).ToNot(HaveOccurred())
		Expect(result.Integer()).To(Equal(42))
	})

	It("Should be able to pass arguments", func() {
		code := `new Function("var args = Array.prototype.slice.call(arguments);return args.join(',')")`
		value, err := state.DoString(code)
		Expect(err).ToNot(HaveOccurred())

		callable, ok := value.(*JsCallable)
		Expect(ok).To(BeTrue())

		result, err := callable.Call(42, "hello", true, 17.89, nil, false)
		Expect(err).ToNot(HaveOccurred())
		Expect(result.String()).To(Equal("42,hello,true,17.89,,false"))
	})

	It("Should properly handle expections", func() {
		value, err := state.DoString(`new Function("throw 'fuuuuu'")`)
		Expect(err).ToNot(HaveOccurred())

		callable, ok := value.(*JsCallable)
		Expect(ok).To(BeTrue())

		_, err = callable.Call()
		Expect(err).To(HaveOccurred())
	})

})
