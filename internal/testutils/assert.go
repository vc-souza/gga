package testutils

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func fail(t *testing.T, v ...any) {
	t.Log(v...)
	t.Log(string(debug.Stack()))
	t.FailNow()
}

// AssertEqual asserts that two values of same type are equal, otherwise the test fails.
func AssertEqual[T comparable](t *testing.T, expected T, actual T) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	fail(t, "expected", expected, "got", actual)
}

// TODO: docs
func AssertTrue(t *testing.T, v bool) {
	if v {
		return
	}

	fail(t, "expected true")
}

// TODO: docs
func AssertFalse(t *testing.T, v bool) {
	if !v {
		return
	}

	fail(t, "expected false")
}

// TODO: docs
func AssertNil(t *testing.T, v any) {
	if reflect.TypeOf(v) == nil {
		return
	}

	if reflect.ValueOf(v).IsNil() {
		return
	}

	fail(t, "expected nil value")
}

// TODO: docs
func AssertNotNil(t *testing.T, v any) {
	if reflect.TypeOf(v) != nil && !reflect.ValueOf(v).IsNil() {
		return
	}

	fail(t, "expected non-nil value")
}
