package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestQueueEnqueue(t *testing.T) {
	var q Queue[int] = new(Deque[int])
	var item int
	var ok bool

	q.Enqueue(1, 2, 3)

	item, ok = q.(*Deque[int]).Get(0)
	ut.Equal(t, true, ok)
	ut.Equal(t, 1, item)

	item, ok = q.(*Deque[int]).Get(1)
	ut.Equal(t, true, ok)
	ut.Equal(t, 2, item)

	item, ok = q.(*Deque[int]).Get(2)
	ut.Equal(t, true, ok)
	ut.Equal(t, 3, item)
}

func TestQueueEmpty(t *testing.T) {
	var q Queue[int] = new(Deque[int])

	ut.Equal(t, true, q.Empty())

	q.Enqueue(1)

	ut.Equal(t, false, q.Empty())
}

func TestQueueDequeue(t *testing.T) {
	var q Queue[int] = new(Deque[int])
	var item int
	var ok bool

	ut.Equal(t, true, q.Empty())

	q.Enqueue(1, 2, 3)

	ut.Equal(t, false, q.Empty())

	item, ok = q.Dequeue()
	ut.Equal(t, true, ok)
	ut.Equal(t, 1, item)

	item, ok = q.Dequeue()
	ut.Equal(t, true, ok)
	ut.Equal(t, 2, item)

	item, ok = q.Dequeue()
	ut.Equal(t, true, ok)
	ut.Equal(t, 3, item)

	ut.Equal(t, true, q.Empty())
}

func TestQueueDequeue_empty(t *testing.T) {
	var q Queue[int] = new(Deque[int])

	ut.Equal(t, true, q.Empty())

	_, ok := q.Dequeue()
	ut.Equal(t, false, ok)
}

func TestQueueDequeue_wrong_type(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("function did not panic")
			t.FailNow()
		}
	}()

	var q Queue[int] = new(Deque[int])

	// forcefully adding an item with wrong type
	if d, ok := q.(*Deque[int]); ok {
		d.PushBack("wrong")
	}

	q.Dequeue()
}
