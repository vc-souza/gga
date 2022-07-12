package ds

// Queue implementations are able to behave like a FIFO (First In / First Out) queue.
type Queue[T any] interface {
	// Enqueue adds an item to the end of the queue.
	Enqueue(...T)

	// Dequeue removes and then returns the item at the start of the queue, if any.
	Dequeue() (T, bool)

	// Empty checks if the queue is empty.
	Empty() bool
}
