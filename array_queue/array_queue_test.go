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
		q := array_queue.New()
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
	q := array_queue.New()
	removedItem, err := q.Dequeue()
	if err == nil {
		t.Errorf("error: dequeuing from empty queue")
	}
	if removedItem != nil {
		t.Errorf("removedItem is not zero valued")
	}
}

func TestIsEmpty(t *testing.T) {
	q := array_queue.New()
	if !q.IsEmpty() {
		t.Errorf("error: queue should be Empty")
	}

	for n := 0; n < 3; n++ {
		q.Enqueue(n)
		if q.IsEmpty() {
			t.Errorf("error: queue should not Empty")
		}
	}

	for n := 0; n < 3; n++ {
		if q.IsEmpty() {
			t.Errorf("error: queue should not Empty")
		}
		q.Dequeue()
	}

	if !q.IsEmpty() {
		t.Errorf("error: queue should be Empty")
	}

}

func TestPeek(t *testing.T) {
	q := array_queue.New()

	for n := 0; n < 3; n++ {
		q.Enqueue(n)
	}

	for n := 0; n < 3; n++ {
		element, _ := q.Peek()
		if element != n {
			t.Errorf("Peek error: got %v expected %v", element, n)
		}
		q.Dequeue()
	}
}

func TestString(t *testing.T) {
	q := array_queue.New()
	for n := 0; n < 3; n++ {
		q.Enqueue(n)
	}
	q.Enqueue("last")

	got := q.String()
	expected := "size = 4 values = [0 1 2 last]"
	if got != expected {
		t.Errorf("String() error: expected\n%s\ngot\n%s\n", expected, got)
	}
}
