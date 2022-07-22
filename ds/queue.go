package ds

import "container/list"

// Queue implementations are able to behave like a FIFO (First In / First Out) queue.
type Queue[T any] interface {
	// Enqueue adds an item to the end of the queue.
	Enqueue(...T)

	// Dequeue removes and then returns the item at the start of the queue, if any.
	Dequeue() (T, bool)

	// Empty checks if the queue is empty.
	Empty() bool
}

// LLQueue implements the Queue interface using a linked list.
type LLQueue[T any] struct {
	list.List
}

// NewQueue returns a new Queue, using an LLQueue implementation.
func NewQueue[T any]() Queue[T] {
	return Queue[T](new(LLQueue[T]))
}

func (d *LLQueue[T]) Enqueue(ts ...T) {
	for _, t := range ts {
		d.PushBack(t)
	}
}

func (d *LLQueue[T]) Dequeue() (T, bool) {
	var zero T

	if d.Empty() {
		return zero, false
	}

	f := d.Front()
	res := d.value(f)

	d.Remove(f)

	return res, true
}

func (d *LLQueue[T]) Empty() bool {
	return d.Len() == 0
}

func (d *LLQueue[T]) get(idx int) (T, bool) {
	var zero T
	count := -1

	for e := d.Front(); e != nil; e = e.Next() {
		count++

		if count != idx {
			continue
		}

		return d.value(e), true
	}

	return zero, false
}

func (d *LLQueue[T]) value(e *list.Element) T {
	if val, ok := e.Value.(T); ok {
		return val
	}

	panic(ErrInvType)
}
