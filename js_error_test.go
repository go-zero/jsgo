package jsgo_test

import (
	"testing"

	. "github.com/go-zero/jsgo"
	. "github.com/onsi/gomega"
	. "github.com/stretchr/testify/suite"
)

type JsErrorTestSuite struct {
	Suite
	state *JsState
}

func (s *JsErrorTestSuite) SetupTest() {
	s.state = NewJsState()
}

func (s *JsErrorTestSuite) TestStandardError() {
	_, err := s.state.DoString("a = 1")
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(ContainSubstring("ReferenceError: assignment to undeclared variable 'a'"))
}

func TestJsErrorTestSuite(t *testing.T) {
	RegisterTestingT(t)
	Run(t, new(JsErrorTestSuite))
}
