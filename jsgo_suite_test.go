package jsgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJSGO(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JSGO Suite")
}
