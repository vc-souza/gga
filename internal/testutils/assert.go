package testutils

import (
	"reflect"
	"runtime/debug"
	"testing"
)

// AssertEQ asserts that two values of same type are equal, otherwise the test fails.
func AssertEQ[T comparable](t *testing.T, expected T, actual T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Log("expected", expected, "got", actual)
		t.Log(string(debug.Stack()))
		t.FailNow()
	}
}
