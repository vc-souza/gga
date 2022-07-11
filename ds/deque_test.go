package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestDequeGet(t *testing.T) {
	d := new(Deque[int])

	d.PushFront(3)

	v, ok := d.Get(0)
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 3, v)
}

func TestDequeGet_invalid(t *testing.T) {
	d := new(Deque[int])

	_, ok := d.Get(-1)
	ut.AssertEqual(t, false, ok)
}

func TestDequeGet_wrong_type(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("function did not panic")
			t.FailNow()
		}
	}()

	d := new(Deque[int])

	d.PushBack("wrong")
	d.Get(0)
}
