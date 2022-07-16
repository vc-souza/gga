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
	ut.AssertEQ(t, true, ok)
	ut.AssertEQ(t, 1, item)

	item, ok = q.(*Deque[int]).Get(1)
	ut.AssertEQ(t, true, ok)
	ut.AssertEQ(t, 2, item)

	item, ok = q.(*Deque[int]).Get(2)
	ut.AssertEQ(t, true, ok)
	ut.AssertEQ(t, 3, item)
}

func TestQueueEmpty(t *testing.T) {
	var q Queue[int] = new(Deque[int])

	ut.AssertEQ(t, true, q.Empty())

	q.Enqueue(1)

	ut.AssertEQ(t, false, q.Empty())
}

func TestQueueDequeue(t *testing.T) {
	var q Queue[int] = new(Deque[int])
	var item int
	var ok bool

	ut.AssertEQ(t, true, q.Empty())

	q.Enqueue(1, 2, 3)

	ut.AssertEQ(t, false, q.Empty())

	item, ok = q.Dequeue()
	ut.AssertEQ(t, true, ok)
	ut.AssertEQ(t, 1, item)

	item, ok = q.Dequeue()
	ut.AssertEQ(t, true, ok)
	ut.AssertEQ(t, 2, item)

	item, ok = q.Dequeue()
	ut.AssertEQ(t, true, ok)
	ut.AssertEQ(t, 3, item)

	ut.AssertEQ(t, true, q.Empty())
}

func TestQueueDequeue_empty(t *testing.T) {
	var q Queue[int] = new(Deque[int])

	ut.AssertEQ(t, true, q.Empty())

	_, ok := q.Dequeue()
	ut.AssertEQ(t, false, ok)
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
