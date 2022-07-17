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

	item, ok = s.(*Deque[int]).Get(2)
	ut.Equal(t, true, ok)
	ut.Equal(t, 1, item)

	item, ok = s.(*Deque[int]).Get(1)
	ut.Equal(t, true, ok)
	ut.Equal(t, 2, item)

	item, ok = s.(*Deque[int]).Get(0)
	ut.Equal(t, true, ok)
	ut.Equal(t, 3, item)
}

func TestStackEmpty(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	ut.Equal(t, true, s.Empty())

	s.Push(1)

	ut.Equal(t, false, s.Empty())
}

func TestStackPeek(t *testing.T) {
	var s Stack[int] = new(Deque[int])
	var item int
	var ok bool

	_, ok = s.Peek()
	ut.Equal(t, false, ok)

	s.Push(1)

	item, ok = s.Peek()
	ut.Equal(t, true, ok)
	ut.Equal(t, 1, item)

	s.Push(2)

	item, ok = s.Peek()
	ut.Equal(t, true, ok)
	ut.Equal(t, 2, item)

	s.Push(3)

	item, ok = s.Peek()
	ut.Equal(t, true, ok)
	ut.Equal(t, 3, item)
}

func TestStackPeek_empty(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	ut.Equal(t, true, s.Empty())

	_, ok := s.Peek()
	ut.Equal(t, false, ok)
}

func TestStackPop(t *testing.T) {
	var s Stack[int] = new(Deque[int])
	var item int
	var ok bool

	ut.Equal(t, true, s.Empty())

	s.Push(1, 2, 3)

	ut.Equal(t, false, s.Empty())

	item, ok = s.Pop()
	ut.Equal(t, true, ok)
	ut.Equal(t, 3, item)

	item, ok = s.Pop()
	ut.Equal(t, true, ok)
	ut.Equal(t, 2, item)

	item, ok = s.Pop()
	ut.Equal(t, true, ok)
	ut.Equal(t, 1, item)

	ut.Equal(t, true, s.Empty())
}

func TestStackPop_empty(t *testing.T) {
	var s Stack[int] = new(Deque[int])

	ut.Equal(t, true, s.Empty())

	_, ok := s.Pop()
	ut.Equal(t, false, ok)
}

func TestStackPop_wrong_type(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("function did not panic")
			t.FailNow()
		}
	}()

	var s Stack[int] = new(Deque[int])

	// forcefully adding an item with wrong type
	if d, ok := s.(*Deque[int]); ok {
		d.PushBack("wrong")
	}

	s.Pop()
}
