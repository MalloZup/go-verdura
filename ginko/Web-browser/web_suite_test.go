package Web_browser_test

// Doc; http://agouti.org/

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
	"testing"
)

func TestWEB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WEB Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	// Choose a WebDriver:

	agoutiDriver = agouti.PhantomJS()
	// agoutiDriver = agouti.Selenium()

	// make sure that chrome driver is in $PATH
	//agoutiDriver = agouti.ChromeDriver()
	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = Describe("UserLogin_Spacewalk", func() {
	var page *agouti.Page
	It("should manage user authentication", func() {
		By("redirecting the user to the login form from the home page", func() {
			page.Navigate("https://www.google.de/?gws_rd=ssl")

		})

	})
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
