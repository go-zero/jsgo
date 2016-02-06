package jsgo_test

import (
	"fmt"

	. "github.com/go-zero/jsgo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JsError", func() {

	var name = "ErrorName"
	var message = "something useful"
	var stackTrace = "\n[foo.js]:1 some issue"

	It("Should generate the expected output", func() {
		err := &JsError{name, message, stackTrace}
		Expect(err.Error()).To(Equal(fmt.Sprintf("%s: %s%s", name, message, stackTrace)))
	})

})
