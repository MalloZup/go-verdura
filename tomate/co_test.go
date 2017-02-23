package cocomero_test

import (
	"../cocomero/"
	"bytes"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// Quick Check generate non-latin string, we want to use only latin strings.
const latin = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01233456789"

type LatinString string

func (l LatinString) Generate(rand *rand.Rand, size int) reflect.Value {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteString(string(latin[rand.Intn(len(latin))]))
	}
	s := LatinString(buffer.String())
	return reflect.ValueOf(s)
}

// We import a package cocomero and use quickcheck

var cucumber = cocomero.Urls{"http://m226.mgr.suse.de/workspace/manager-Head-sumaform-cucumber/last.html",
	"http://m226.mgr.suse.de/workspace/manager-3.0-sumaform30/last.html",
	"last.html",
	"last30.html"}

func TestCheck(t *testing.T) {

	// this example show how to test a property of a function
	// we generate random latin string and integer x

	// we want to test that data is always string type, after we call ReadLastResults
	assertion := func(head LatinString, x int) bool {
		// make your assertion here.
		var s string
		typ := reflect.TypeOf(s)
		cucumber.GetHeadOutput()
		data := cucumber.ReadLastResults()
		return reflect.TypeOf(data) == typ
	}
	t.Log("run quick check test")
	// Run the assertion through the quick checker
	if err := quick.Check(assertion, nil); err != nil {
		t.Log(err)
	}
}
