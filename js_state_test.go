package jsgo_test

import (
	"runtime"
	"testing"

	. "github.com/go-zero/jsgo"
	. "github.com/onsi/gomega"
	. "github.com/stretchr/testify/suite"
)

type JsStateTestSuite struct {
	Suite
	state *JsState
}

func (s *JsStateTestSuite) SetupTest() {
	s.state = NewJsState()
}

func (s *JsStateTestSuite) TearDownTest() {
	runtime.GC()
}

func (s *JsStateTestSuite) TestToNotBeNil() {
	Expect(s.state).ToNot(BeNil())
}

func (s *JsStateTestSuite) TestRunAPieceOfCode() {
	_, err := s.state.DoString("1")
	Expect(err).ToNot(HaveOccurred())
}

func (s *JsStateTestSuite) TestErrorHandling() {
	_, err := s.state.DoString("a")
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(ContainSubstring("ReferenceError"))
}

func (s *JsStateTestSuite) TestSyntaxErrorHandling() {
	_, err := s.state.DoString("1{}")
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(ContainSubstring("SyntaxError"))
}

func TestJsStateTestSuite(t *testing.T) {
	RegisterTestingT(t)
	Run(t, new(JsStateTestSuite))
}
