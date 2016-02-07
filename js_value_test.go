package jsgo_test

import (
	"testing"

	. "github.com/go-zero/jsgo"
	. "github.com/onsi/gomega"
	. "github.com/stretchr/testify/suite"
)

type CheckFunction func(JsValue) bool

type JsValueTestSuite struct {
	Suite
	state *JsState
}

func (s *JsValueTestSuite) SetupTest() {
	s.state = NewJsState()
}

func (s *JsValueTestSuite) TestTypes() {
	checks := map[string]CheckFunction{
		"1":              JsValue.IsDefined,
		"":               JsValue.IsUndefined,
		"42.0":           JsValue.IsNumber,
		"'hola'":         JsValue.IsString,
		"true":           JsValue.IsBool,
		"null":           JsValue.IsNull,
		"false":          JsValue.IsPrimitive,
		"new Function()": JsValue.IsCallable,
		"new Array()":    JsValue.IsArray,
		"new RegExp()":   JsValue.IsRegExp,
		"new Object()":   JsValue.IsObject,
	}

	for code, check := range checks {
		value, err := s.state.DoString(code)
		Expect(err).ToNot(HaveOccurred())
		Expect(check(value)).To(BeTrue())
	}
}

func (s *JsValueTestSuite) TestReturnAFloat() {
	value, err := s.state.DoString("182.0")
	Expect(err).ToNot(HaveOccurred())
	Expect(value.Float()).To(Equal(182.0))
}

func (s *JsValueTestSuite) TestReturnAnInteger() {
	value, err := s.state.DoString("42")
	Expect(err).ToNot(HaveOccurred())
	Expect(value.Integer()).To(Equal(42))
}

func (s *JsValueTestSuite) TestReturnAString() {
	value, err := s.state.DoString("'Hello World!'")
	Expect(err).ToNot(HaveOccurred())
	Expect(value.String()).To(Equal("Hello World!"))
}

func (s *JsValueTestSuite) TestReturnABoolean() {
	value, err := s.state.DoString("true")
	Expect(err).ToNot(HaveOccurred())
	Expect(value.Bool()).To(BeTrue())
}

func (s *JsValueTestSuite) TestReturnACallable() {
	value, err := s.state.DoString("new Function()")
	Expect(err).ToNot(HaveOccurred())
	_, ok := value.(*JsCallable)
	Expect(ok).To(BeTrue())
}

func (s *JsValueTestSuite) TestReturnAnObject() {
	value, err := s.state.DoString("new Object()")
	Expect(err).ToNot(HaveOccurred())
	_, ok := value.(*JsObject)
	Expect(ok).To(BeTrue())
}

func (s *JsValueTestSuite) TestReturnAnArray() {
	value, err := s.state.DoString("new Array()")
	Expect(err).ToNot(HaveOccurred())
	_, ok := value.(*JsArray)
	Expect(ok).To(BeTrue())
}

func (s *JsValueTestSuite) TestReturnARegExp() {
	value, err := s.state.DoString("new RegExp()")
	Expect(err).ToNot(HaveOccurred())
	_, ok := value.(*JsRegExp)
	Expect(ok).To(BeTrue())
}

func TestJsValueTestSuite(t *testing.T) {
	RegisterTestingT(t)
	Run(t, new(JsValueTestSuite))
}
