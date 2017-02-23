package cocomero_test

import (
	"../cocomero/"
	"reflect"
	"testing"
	"testing/quick"
)

// We import a package cocomero and use quickcheck

var cucumber = cocomero.Urls{"http://m226.mgr.suse.de/workspace/manager-Head-sumaform-cucumber/last.html",
	"http://m226.mgr.suse.de/workspace/manager-3.0-sumaform30/last.html",
	"last.html",
	"last30.html"}

func TestCheck(t *testing.T) {

	assertion := func(head string, x int) bool {
		// make your assertion here.
		typ := reflect.TypeOf(head)
		t.Log(typ)
		t.Logf("random head %s", string(head))
		t.Logf("random int %v", x)
		cucumber.GetHeadOutput()

		return cucumber.HeadFile == "last.html"
	}
	t.Log("run quick check test")
	// Run the assertion through the quick checker
	if err := quick.Check(assertion, nil); err != nil {
		t.Log(err)
	}
}
