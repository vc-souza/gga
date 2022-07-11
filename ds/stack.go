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
