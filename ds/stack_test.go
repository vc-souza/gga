package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestStackPush(t *testing.T) {
	var s Stack[int] = new(Deque[int])
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

func TestStackEmpty(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	ut.AssertEqual(t, true, s.Empty())

	s.Push(1)

	ut.AssertEqual(t, false, s.Empty())
}

func TestStackPeek(t *testing.T) {
	var s Stack[int] = new(Deque[int])
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

func TestStackPeek_empty(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	ut.AssertEqual(t, true, s.Empty())

	_, ok := s.Peek()
	ut.AssertEqual(t, false, ok)
}

func TestStackPop(t *testing.T) {
	var s Stack[int] = new(Deque[int])
	var item int
	var ok bool

	ut.AssertEqual(t, true, s.Empty())

	s.Push(1, 2, 3)

	ut.AssertEqual(t, false, s.Empty())

	item, ok = s.Pop()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 3, item)

	item, ok = s.Pop()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 2, item)

	item, ok = s.Pop()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 1, item)

	ut.AssertEqual(t, true, s.Empty())
}

func TestStackPop_empty(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	ut.AssertEqual(t, true, s.Empty())

	_, ok := s.Pop()
	ut.AssertEqual(t, false, ok)
}

func TestStackPop_wrong_type(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	// forcefully adding an item with wrong type
	if d, ok := s.(*Deque[int]); ok {
		d.PushBack("wrong")
	}

	_, ok := s.Pop()
	ut.AssertEqual(t, false, ok)
}

func TestStackGet_invalid(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	_, ok := s.Get(-1)
	ut.AssertEqual(t, false, ok)
}

func TestStackGet_wrong_type(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	// forcefully adding an item with wrong type
	if d, ok := s.(*Deque[int]); ok {
		d.PushBack("wrong")
	}

	_, ok := s.Get(0)
	ut.AssertEqual(t, false, ok)
}
