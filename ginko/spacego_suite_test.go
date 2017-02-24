package zenzero_test

import (
	"./features/"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestFeatures(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testsuite Spacego ")
}

var _ = Describe("Testsuite Spacego", func() {
	features.Zenzero()
	features.Broccoli()
	features.Spinaci()
})
