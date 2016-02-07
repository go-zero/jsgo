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

	It("Should return a float value", func() {
		value, _ := state.DoString("182.0")
		Expect(value.Float()).To(Equal(182.0))
	})

	It("Should return an integer value", func() {
		value, _ := state.DoString("42")
		Expect(value.Integer()).To(Equal(42))
	})

	It("Should return a string value", func() {
		value, _ := state.DoString("'Hello World!'")
		Expect(value.String()).To(Equal("Hello World!"))
	})

	It("Should return a boolean", func() {
		value, _ := state.DoString("true")
		Expect(value.Bool()).To(Equal(true))
	})

	It("Should return a callable", func() {
		value, _ := state.DoString("new Function()")
		_, ok := value.(*JsCallable)
		Expect(ok).To(Equal(true))
	})

	It("Should return an object", func() {
		value, _ := state.DoString("new Object()")
		_, ok := value.(*JsObject)
		Expect(ok).To(Equal(true))
	})

	It("Should return an array", func() {
		value, _ := state.DoString("new Array()")
		_, ok := value.(*JsArray)
		Expect(ok).To(Equal(true))
	})

	It("Should return a regexp", func() {
		value, _ := state.DoString("new RegExp()")
		_, ok := value.(*JsRegExp)
		Expect(ok).To(Equal(true))
	})

	DescribeTable("Should identify the value type",
		func(code string, check CheckFunction) {
			value, _ := state.DoString(code)
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
