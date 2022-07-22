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

func Equal[T comparable](t *testing.T, expected T, actual T) {
	if reflect.DeepEqual(expected, actual) {
		return
	}

	fail(t, "expected", expected, "got", actual)
}

func True(t *testing.T, v bool) {
	if v {
		return
	}

	fail(t, "expected true")
}

func False(t *testing.T, v bool) {
	if !v {
		return
	}

	fail(t, "expected false")
}

func Nil(t *testing.T, v any) {
	if reflect.TypeOf(v) == nil {
		return
	}

	if reflect.ValueOf(v).IsNil() {
		return
	}

	fail(t, "expected nil value")
}

func NotNil(t *testing.T, v any) {
	if reflect.TypeOf(v) != nil && !reflect.ValueOf(v).IsNil() {
		return
	}

	fail(t, "expected non-nil value")
}
