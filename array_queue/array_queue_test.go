package array_queue_test

import (
	"github.com/bnixon67/go_queue/array_queue"
	"testing"
)

func TestEnqueueDequeue(t *testing.T) {
	tests := []struct {
		data []interface{}
	}{
		{[]interface{}{42}},
		{[]interface{}{"one", 2, "three", 4, "five"}},
	}

	for _, test := range tests {
		q := array_queue.NewQueue()
		for i, element := range test.data {
			q.Enqueue(element)
			if i+1 != q.Size() {
				t.Errorf("Incorrect Size expected: %d got: %d",
					i+1, q.Size())
			}
		}

		for i := 0; i < len(test.data); i++ {
			removedItem, err := q.Dequeue()
			if err != nil {
				t.Errorf("error while dequeue")
			}
			if removedItem != test.data[i] {
				t.Errorf("Wrong output expected: %d got: %d",
					test.data[i], removedItem)
			}
		}
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := array_queue.NewQueue()
	removedItem, err := q.Dequeue()
	if err == nil {
		t.Errorf("error: dequeuing from empty queue")
	}
	if removedItem != nil {
		t.Errorf("removedItem is not zero valued")
	}
}
