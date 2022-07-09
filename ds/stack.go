package ds

// Stack implementations are able to behave like a LIFO (Last In - First Out) stack.
type Stack[T any] interface {
	// Peek returns the item at the top of the stack, if any.
	Peek() (T, bool)

	// Pop removes and then returns the item at the top of the stack, if any.
	Pop() (T, bool)

	// Push adds an item at the top of the stack.
	Push(...T)

	// Empty checks if the stack is empty.
	Empty() bool

	// Get fetches an item at a particular position, for testability
	Get(int) (T, bool)
}

// SliceStack is a slice implementation of the Stack interface.
type SliceStack[T any] []T

func (s SliceStack[T]) Peek() (T, bool) {
	var res T

	if len(s) == 0 {
		return res, false
	}

	return s[len(s)-1], true
}

func (s *SliceStack[T]) Pop() (T, bool) {
	var zero T
	var t T

	if len(*s) == 0 {
		return zero, false
	}

	// get a reference / copy of the last item, the top of the stack
	t = (*s)[len(*s)-1]

	// set the last item to the zero value of its type
	// for pointer types, this avoids a memory leak
	(*s)[len(*s)-1] = zero

	// reassigns the slice after slicing out the last item
	*s = (*s)[:len(*s)-1]

	return t, true
}

func (s *SliceStack[T]) Push(t ...T) {
	*s = append(*s, t...)
}

func (s *SliceStack[T]) Empty() bool {
	return len(*s) == 0
}

func (s *SliceStack[T]) Get(idx int) (T, bool) {
	var res T

	if idx < 0 || idx >= len(*s) {
		return res, false
	}

	return (*s)[idx], true
}
