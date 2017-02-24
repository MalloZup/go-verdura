package zenzero_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestZenzero(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Zenzero Suite")
}

var _ = Describe("Zenzero", func() {

	Context("initially", func() {
		It("has 0 items", func() {})
		It("has 0 units", func() {})
		Specify("the total amount is 0.00", func() {})
	})
	Context("when a new item is added", func() {
		Context("the shopping cart", func() {
			It("has 1 more unique item than it had earlier", func() {})
			It("has 1 more unit than it had earlier", func() {})
			Specify("total amount increases by item price", func() {})
		})
	})
})
