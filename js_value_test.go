package jsgo_test

import (
	"runtime"

	. "github.com/go-zero/jsgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

type CheckFunction func(JsValue) bool

var _ = Describe("JsValue", func() {
	state := NewJsState()

	It("Should return a float", func() {
		value, err := state.DoString("182.0")
		Expect(err).ToNot(HaveOccurred())
		Expect(value.Float()).To(Equal(182.0))
	})

	It("Should return an integer", func() {
		value, err := state.DoString("42")
		Expect(err).ToNot(HaveOccurred())
		Expect(value.Integer()).To(Equal(42))
	})

	It("Should return a string", func() {
		value, err := state.DoString("'Hello World!'")
		Expect(err).ToNot(HaveOccurred())
		Expect(value.String()).To(Equal("Hello World!"))
	})

	It("Should return a boolean", func() {
		value, err := state.DoString("true")
		Expect(err).ToNot(HaveOccurred())
		Expect(value.Bool()).To(BeTrue())
	})

	It("Should return a callable", func() {
		value, err := state.DoString("new Function()")
		Expect(err).ToNot(HaveOccurred())
		_, ok := value.(*JsCallable)
		Expect(ok).To(BeTrue())
	})

	It("Should return an object", func() {
		value, err := state.DoString("new Object()")
		Expect(err).ToNot(HaveOccurred())
		_, ok := value.(*JsObject)
		Expect(ok).To(BeTrue())
	})

	It("Should return an array", func() {
		value, err := state.DoString("new Array()")
		Expect(err).ToNot(HaveOccurred())
		_, ok := value.(*JsArray)
		Expect(ok).To(BeTrue())
	})

	It("Should return a regexp", func() {
		value, err := state.DoString("new RegExp()")
		Expect(err).ToNot(HaveOccurred())
		_, ok := value.(*JsRegExp)
		Expect(ok).To(BeTrue())
	})

	DescribeTable("Should identify the value type",
		func(code string, check CheckFunction) {
			value, err := state.DoString(code)
			Expect(err).ToNot(HaveOccurred())
			Expect(check(value)).To(BeTrue())
			runtime.GC() // force the unreferencing of the value
		},
		Entry("defined", "1", JsValue.IsDefined),
		Entry("undefined", "{}", JsValue.IsUndefined),
		Entry("number", "1", JsValue.IsNumber),
		Entry("string", "'hola'", JsValue.IsString),
		Entry("boolean", "true", JsValue.IsBool),
		Entry("null", "null", JsValue.IsNull),
		Entry("primitive", "false", JsValue.IsPrimitive),
		Entry("callable", "new Function()", JsValue.IsCallable),
		Entry("array", "new Array()", JsValue.IsArray),
		Entry("regexp", "new RegExp()", JsValue.IsRegExp),
		Entry("object", "new Object()", JsValue.IsObject),
	)

})
