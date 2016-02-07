package jsgo_test

import (
	"testing"

	. "github.com/go-zero/jsgo"
	. "github.com/onsi/gomega"
	. "github.com/stretchr/testify/suite"
)

type JsCallableTestSuite struct {
	Suite
	state *JsState
}

func (s *JsCallableTestSuite) SetupTest() {
	s.state = NewJsState()
}

func (s *JsCallableTestSuite) TestBeCallable() {
	value, err := s.state.DoString("new Function('return 42')")
	Expect(err).ToNot(HaveOccurred())

	callable, ok := value.(*JsCallable)
	Expect(ok).To(BeTrue())

	result, err := callable.Call()
	Expect(err).ToNot(HaveOccurred())
	Expect(result.Integer()).To(Equal(42))
}

func (s *JsCallableTestSuite) TestCallWithMultipleArgs() {
	code := `new Function("var args = Array.prototype.slice.call(arguments);return args.join(',')")`
	value, err := s.state.DoString(code)
	Expect(err).ToNot(HaveOccurred())

	callable, ok := value.(*JsCallable)
	Expect(ok).To(BeTrue())

	result, err := callable.Call(42, "hello", true, 17.89, nil, false)
	Expect(err).ToNot(HaveOccurred())
	Expect(result.String()).To(Equal("42,hello,true,17.89,,false"))
}

func (s *JsCallableTestSuite) TestToReturnError() {
	value, err := s.state.DoString(`new Function("throw 'fuuuuu'")`)
	Expect(err).ToNot(HaveOccurred())

	callable, ok := value.(*JsCallable)
	Expect(ok).To(BeTrue())

	_, err = callable.Call()
	Expect(err).To(HaveOccurred())
}

func TestJsCallableTestSuite(t *testing.T) {
	RegisterTestingT(t)
	Run(t, new(JsCallableTestSuite))
}
