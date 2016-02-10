package jsgo_test

import (
	"runtime"
	"testing"

	. "github.com/go-zero/jsgo"
	. "github.com/onsi/gomega"
	. "github.com/stretchr/testify/suite"
)

type JsGlobalTestSuite struct {
	Suite
	state *JsState
}

func (s *JsGlobalTestSuite) SetupTest() {
	s.state = NewJsState()
}

func (s *JsGlobalTestSuite) TearDownTest() {
	runtime.GC()
}

func (s *JsGlobalTestSuite) TestGetAndSetGlobalVariable() {
	value, err := s.state.DoString("42")
	Expect(err).ToNot(HaveOccurred())

	s.state.Global().Set("answer", value)
	Expect(s.state.Global().Get("answer").Integer()).To(Equal(42))
}

func TestJsGlobalTestSuite(t *testing.T) {
	RegisterTestingT(t)
	Run(t, new(JsGlobalTestSuite))
}
