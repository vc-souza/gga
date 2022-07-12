package ds

import "container/list"

/*
A Deque (Double-Ended Queue) allows insertion and removal at both ends of a doubly linked list,
which is already provided by list.List, from the standard library. So the idea is to embed a
list.List, and use it to implement both the ds.Queue and ds.Stack interfaces.
*/
type Deque[T any] struct {
	list.List
}

func (d *Deque[T]) Enqueue(ts ...T) {
	for _, t := range ts {
		d.PushBack(t)
	}
}

func (d *Deque[T]) Dequeue() (T, bool) {
	return d.top(true)
}

func (d *Deque[T]) Peek() (T, bool) {
	return d.top(false)
}

func (d *Deque[T]) Pop() (T, bool) {
	return d.top(true)
}

func (d *Deque[T]) Push(ts ...T) {
	for _, t := range ts {
		d.PushFront(t)
	}
}

func (d *Deque[T]) Empty() bool {
	return d.Len() == 0
}

func (d *Deque[T]) Get(idx int) (T, bool) {
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

func (d *Deque[T]) value(e *list.Element) T {
	if val, ok := e.Value.(T); ok {
		return val
	}

	panic(ErrInvalidType)
}

func (d *Deque[T]) top(rm bool) (T, bool) {
	var zero T

	if d.Empty() {
		return zero, false
	}

	f := d.Front()
	res := d.value(f)

	if rm {
		d.Remove(f)
	}

	return res, true
}
