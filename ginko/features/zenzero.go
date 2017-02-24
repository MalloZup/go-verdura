package features

// in steps directory we just make the functions
import (
	"../steps/"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

func Zenzero() {

	var _ = Describe("Zenzero", func() {

		Context("Try a login", func() {
			It("I use the function login from auth package", func() { auth.Login() })
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
}
