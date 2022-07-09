package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestQueueEnqueue(t *testing.T) {
	var q Queue[int] = new(LLQueue[int])
	var item int
	var ok bool

	q.Enqueue(1, 2, 3)

	item, ok = q.Get(0)
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 1, item)

	item, ok = q.Get(1)
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 2, item)

	item, ok = q.Get(2)
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 3, item)
}

func TestQueueEmpty(t *testing.T) {
	var q Queue[int] = new(LLQueue[int])

	ut.AssertEqual(t, true, q.Empty())

	q.Enqueue(1)

	ut.AssertEqual(t, false, q.Empty())
}

func TestQueueDequeue(t *testing.T) {
	var q Queue[int] = new(LLQueue[int])
	var item int
	var ok bool

	ut.AssertEqual(t, true, q.Empty())

	q.Enqueue(1, 2, 3)

	ut.AssertEqual(t, false, q.Empty())

	item, ok = q.Dequeue()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 1, item)

	item, ok = q.Dequeue()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 2, item)

	item, ok = q.Dequeue()
	ut.AssertEqual(t, true, ok)
	ut.AssertEqual(t, 3, item)

	ut.AssertEqual(t, true, q.Empty())
}

func TestQueueDequeue_empty(t *testing.T) {
	var q Queue[int] = new(LLQueue[int])

	ut.AssertEqual(t, true, q.Empty())

	_, ok := q.Dequeue()
	ut.AssertEqual(t, false, ok)
}

func TestQueueDequeue_wrong_type(t *testing.T) {
	var q Queue[int] = new(LLQueue[int])

	// forcefully adding an item with wrong type
	if llq, ok := q.(*LLQueue[int]); ok {
		llq.PushBack("wrong")
	}

	_, ok := q.Dequeue()
	ut.AssertEqual(t, false, ok)
}

func TestQueueGet_invalid(t *testing.T) {
	var q Queue[int] = new(LLQueue[int])

	_, ok := q.Get(0)
	ut.AssertEqual(t, false, ok)
}

func TestQueueGet_wrong_type(t *testing.T) {
	var q Queue[int] = new(LLQueue[int])

	// forcefully adding an item with wrong type
	if llq, ok := q.(*LLQueue[int]); ok {
		llq.PushBack("wrong")
	}

	_, ok := q.Get(0)
	ut.AssertEqual(t, false, ok)
}
