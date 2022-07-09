package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestPush(t *testing.T) {
	var s Stack[int] = &SliceStack[int]{}
	var item int
	var ok bool

	s.Push(1, 2, 3)

	item, ok = s.Get(0)
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 1, item)

	item, ok = s.Get(1)
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 2, item)

	item, ok = s.Get(2)
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 3, item)
}

func TestEmpty(t *testing.T) {
	var s Stack[int] = &SliceStack[int]{}

	ut.AssertEqual(t, true, s.Empty())

	s.Push(1)

	ut.AssertEqual(t, false, s.Empty())
}

func TestPeek(t *testing.T) {
	var s Stack[int] = &SliceStack[int]{}
	var item int
	var ok bool

	_, ok = s.Peek()
	ut.AssertEqual(t, false, ok)

	s.Push(1)

	item, ok = s.Peek()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 1, item)

	s.Push(2)

	item, ok = s.Peek()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 2, item)

	s.Push(3)

	item, ok = s.Peek()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 3, item)
}

func TestPop(t *testing.T) {
	var s Stack[int] = &SliceStack[int]{}
	var item int
	var ok bool

	ut.AssertEqual(t, true, s.Empty())

	s.Push(1)

	ut.AssertEqual(t, false, s.Empty())

	item, ok = s.Pop()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 1, item)

	ut.AssertEqual(t, true, s.Empty())
}
