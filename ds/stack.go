package ds

// Stack implementations are able to behave like a LIFO (Last In / First Out) stack.
type Stack[T any] interface {
	// Push adds an item at the top of the stack.
	Push(...T)

	// Peek returns the item at the top of the stack, if any.
	Peek() (T, bool)

	// Pop removes and then returns the item at the top of the stack, if any.
	Pop() (T, bool)

	// Empty checks if the stack is empty.
	Empty() bool
}

// SliceStack implements the Stack interface using a slice.
type SliceStack[T any] []T

// NewStack returns a new Stack, using a SliceStack implementation.
func NewStack[T any]() Stack[T] {
	data := SliceStack[T]([]T{})
	return Stack[T](&data)
}

func (s *SliceStack[T]) Push(ts ...T) {
	for _, t := range ts {
		*s = append(*s, t)
	}
}

func (s SliceStack[T]) Peek() (T, bool) {
	var zero T

	if len(s) == 0 {
		return zero, false
	}

	return s[len(s)-1], true
}

func (s *SliceStack[T]) Pop() (T, bool) {
	var zero T

	if len(*s) == 0 {
		return zero, false
	}

	// get top value
	res := (*s)[len(*s)-1]

	// store zero value at
	// the top position
	(*s)[len(*s)-1] = zero

	// shrink slice, removing
	//
	*s = (*s)[:len(*s)-1]

	return res, true
}

func (s SliceStack[T]) Empty() bool {
	return len(s) == 0
}
