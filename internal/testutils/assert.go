package testutils

import (
	"reflect"
	"testing"
)

// AssertEqual asserts that two values of same type are equal, otherwise the test fails.
func AssertEqual[T comparable](t *testing.T, expected T, actual T) {
	if !reflect.DeepEqual(expected, actual) {
		t.Log("expected", expected, "got", actual)
		t.FailNow()
	}
}
