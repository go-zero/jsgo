package jsgo_test

import (
	"runtime"

	. "github.com/go-zero/jsgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

type CheckFunction func(*JsValue) bool

var _ = Describe("JsValue", func() {
	state := NewJsState()

	It("Should return the string representation", func() {
		value, _ := state.DoString("'Hello World!'")
		Expect(value.String()).To(Equal("Hello World!"))
	})

	DescribeTable("Should identify the value type",
		func(code string, check CheckFunction) {
			value, _ := state.DoString(code)
			Expect(check(value)).To(BeTrue())
			runtime.GC() // force the unreferencing of the value
		},
		Entry("defined", "1", (*JsValue).IsDefined),
		Entry("undefined", "{}", (*JsValue).IsUndefined),
		Entry("number", "1", (*JsValue).IsNumber),
		Entry("string", "'hola'", (*JsValue).IsString),
		Entry("boolean", "true", (*JsValue).IsBoolean),
		Entry("null", "null", (*JsValue).IsNull),
		Entry("primitive", "false", (*JsValue).IsPrimitive),
		Entry("callable", "new Function()", (*JsValue).IsCallable),
		Entry("object", "new Object()", (*JsValue).IsObject),
		Entry("array", "new Array()", (*JsValue).IsArray),
		Entry("regexp", "new RegExp()", (*JsValue).IsRegexp),
	)

})
