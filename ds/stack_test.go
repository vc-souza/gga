package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestStackPush(t *testing.T) {
	var item int

	s := NewStack[int]()

	s.Push(1, 2, 3)

	item = (*s.(*SliceStack[int]))[2]
	ut.AssertEqual(t, 3, item)

	item = (*s.(*SliceStack[int]))[1]
	ut.AssertEqual(t, 2, item)

	item = (*s.(*SliceStack[int]))[0]
	ut.AssertEqual(t, 1, item)
}

func TestStackEmpty(t *testing.T) {
	s := NewStack[int]()

	ut.AssertTrue(t, s.Empty())

	s.Push(1)

	ut.AssertFalse(t, s.Empty())
}

func TestStackPeek(t *testing.T) {
	var item int
	var ok bool

	s := NewStack[int]()

	_, ok = s.Peek()
	ut.AssertFalse(t, ok)

	s.Push(1)

	item, ok = s.Peek()
	ut.AssertTrue(t, ok)
	ut.AssertEqual(t, 1, item)

	s.Push(2)

	item, ok = s.Peek()
	ut.AssertTrue(t, ok)
	ut.AssertEqual(t, 2, item)

	s.Push(3)

	item, ok = s.Peek()
	ut.AssertTrue(t, ok)
	ut.AssertEqual(t, 3, item)
}

func TestStackPeek_empty(t *testing.T) {
	s := NewStack[int]()

	ut.AssertTrue(t, s.Empty())

	_, ok := s.Peek()
	ut.AssertFalse(t, ok)
}

func TestStackPop(t *testing.T) {
	var item int
	var ok bool

	s := NewStack[int]()

	ut.AssertTrue(t, s.Empty())

	s.Push(1, 2, 3)

	ut.AssertFalse(t, s.Empty())

	item, ok = s.Pop()
	ut.AssertTrue(t, ok)
	ut.AssertEqual(t, 3, item)

	item, ok = s.Pop()
	ut.AssertTrue(t, ok)
	ut.AssertEqual(t, 2, item)

	item, ok = s.Pop()
	ut.AssertTrue(t, ok)
	ut.AssertEqual(t, 1, item)

	ut.AssertTrue(t, s.Empty())
}

func TestStackPop_empty(t *testing.T) {
	s := NewStack[int]()

	ut.AssertTrue(t, s.Empty())

	_, ok := s.Pop()
	ut.AssertFalse(t, ok)
}
