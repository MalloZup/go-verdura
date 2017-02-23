package quicktemplate

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func TestValue(t *testing.T) {

	// https://golang.org/pkg/testing/quick/#Check
	// Value returns an arbitrary value of the given type
	var x int

	t.Log("run quick check VALUE test")
	typ := reflect.TypeOf(x)
	var rnd = rand.New(rand.NewSource(42))
	t.Log(quick.Value(typ, rnd))
}

// quick.Check function
func TestCheck(t *testing.T) {
	/*
		Create an assertion, which is a function that takes N inputs of
		random data and returns true if the assertion passes.

		In this case, we're saying take any random integer (x)
		If you add 0, it should equal x
	*/
	assertion := func(x int) bool {
		// make your assertion here.
		return x+0 == x
	}
	t.Log("run quick check test")
	// Run the assertion through the quick checker
	if err := quick.Check(assertion, nil); err != nil {
		fmt.Println("error")
	}
}
