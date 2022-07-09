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

	// Get fetches an item at a particular position, for testability purposes.
	Get(int) (T, bool)
}

// LLQueue is a doubly-linked list implementation of the Queue interface.
type LLQueue[T any] struct {
	list.List
}

func (q *LLQueue[T]) Enqueue(ts ...T) {
	for _, t := range ts {
		q.PushBack(t)
	}
}

func (q *LLQueue[T]) Dequeue() (T, bool) {
	var zero T

	if q.Empty() {
		return zero, false
	}

	f := q.Front()

	if res, ok := f.Value.(T); ok {
		q.Remove(f)
		return res, true
	}

	return zero, false
}

func (q *LLQueue[T]) Empty() bool {
	return q.Len() == 0
}

func (q *LLQueue[T]) Get(idx int) (T, bool) {
	var zero T
	count := -1

	for e := q.Front(); e != nil; e = e.Next() {
		count++

		if count != idx {
			continue
		}

		if res, ok := e.Value.(T); ok {
			return res, true
		}

		break
	}

	return zero, false
}
