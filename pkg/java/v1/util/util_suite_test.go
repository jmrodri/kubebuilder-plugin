package util

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "util")
}

var _ = BeforeSuite(func() {
	os.Setenv("FOO_BAR", "1")
})
