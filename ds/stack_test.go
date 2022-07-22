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
	ut.Equal(t, 3, item)

	item = (*s.(*SliceStack[int]))[1]
	ut.Equal(t, 2, item)

	item = (*s.(*SliceStack[int]))[0]
	ut.Equal(t, 1, item)
}

func TestStackEmpty(t *testing.T) {
	s := NewStack[int]()

	ut.True(t, s.Empty())

	s.Push(1)

	ut.False(t, s.Empty())
}

func TestStackPeek(t *testing.T) {
	var item int
	var ok bool

	s := NewStack[int]()

	_, ok = s.Peek()
	ut.False(t, ok)

	s.Push(1)

	item, ok = s.Peek()
	ut.True(t, ok)
	ut.Equal(t, 1, item)

	s.Push(2)

	item, ok = s.Peek()
	ut.True(t, ok)
	ut.Equal(t, 2, item)

	s.Push(3)

	item, ok = s.Peek()
	ut.True(t, ok)
	ut.Equal(t, 3, item)
}

func TestStackPeek_empty(t *testing.T) {
	s := NewStack[int]()

	ut.True(t, s.Empty())

	_, ok := s.Peek()
	ut.False(t, ok)
}

func TestStackPop(t *testing.T) {
	var item int
	var ok bool

	s := NewStack[int]()

	ut.True(t, s.Empty())

	s.Push(1, 2, 3)

	ut.False(t, s.Empty())

	item, ok = s.Pop()
	ut.True(t, ok)
	ut.Equal(t, 3, item)

	item, ok = s.Pop()
	ut.True(t, ok)
	ut.Equal(t, 2, item)

	item, ok = s.Pop()
	ut.True(t, ok)
	ut.Equal(t, 1, item)

	ut.True(t, s.Empty())
}

func TestStackPop_empty(t *testing.T) {
	s := NewStack[int]()

	ut.True(t, s.Empty())

	_, ok := s.Pop()
	ut.False(t, ok)
}
