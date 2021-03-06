package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestQueueEnqueue(t *testing.T) {
	var item int
	var ok bool

	q := NewQueue[int]()

	q.Enqueue(1, 2, 3)

	item, ok = q.(*LLQueue[int]).get(0)
	ut.True(t, ok)
	ut.Equal(t, 1, item)

	item, ok = q.(*LLQueue[int]).get(1)
	ut.True(t, ok)
	ut.Equal(t, 2, item)

	item, ok = q.(*LLQueue[int]).get(2)
	ut.True(t, ok)
	ut.Equal(t, 3, item)
}

func TestQueueEmpty(t *testing.T) {
	q := NewQueue[int]()

	ut.True(t, q.Empty())

	q.Enqueue(1)

	ut.False(t, q.Empty())
}

func TestQueueDequeue(t *testing.T) {
	var item int
	var ok bool

	q := NewQueue[int]()

	ut.True(t, q.Empty())

	q.Enqueue(1, 2, 3)

	ut.False(t, q.Empty())

	item, ok = q.Dequeue()
	ut.True(t, ok)
	ut.Equal(t, 1, item)

	item, ok = q.Dequeue()
	ut.True(t, ok)
	ut.Equal(t, 2, item)

	item, ok = q.Dequeue()
	ut.True(t, ok)
	ut.Equal(t, 3, item)

	ut.True(t, q.Empty())
}

func TestQueueDequeue_empty(t *testing.T) {
	q := NewQueue[int]()

	ut.True(t, q.Empty())

	_, ok := q.Dequeue()
	ut.False(t, ok)
}

func TestQueueDequeue_wrong_type(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("function did not panic")
			t.FailNow()
		}
	}()

	q := NewQueue[int]()

	// forcefully adding an item with wrong type
	if d, ok := q.(*LLQueue[int]); ok {
		d.PushBack("wrong")
	}

	q.Dequeue()
}

func TestLLQueueGet(t *testing.T) {
	d := new(LLQueue[int])

	d.PushFront(3)

	v, ok := d.get(0)
	ut.True(t, ok)
	ut.Equal(t, 3, v)
}

func TestLLQueueGet_invalid(t *testing.T) {
	d := new(LLQueue[int])

	_, ok := d.get(-1)
	ut.False(t, ok)
}

func TestLLQueueGet_wrong_type(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("function did not panic")
			t.FailNow()
		}
	}()

	d := new(LLQueue[int])

	d.PushBack("wrong")
	d.get(0)
}
