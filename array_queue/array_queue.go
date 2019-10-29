// Package array_queue implements a queue backed by an array.
// A queue is a First In-First Out (FIFO) data type, i.e. the element enqueued
// first will be dequeued first.
package array_queue

import (
	"errors"
	"fmt"
)

// Queue represents a queue backed by an array.
type Queue struct {
	array []interface{}
}

// New returns an initialized empty queue.
func New() *Queue {
	return &Queue{array: make([]interface{}, 0)}
}

// IsEmpty reports whether the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.Size() <= 0
}

// Size reports the number of elements are in the queue.
func (q *Queue) Size() int {
	return len(q.array)
}

// Enqueue adds the element to the end of the queue.
func (q *Queue) Enqueue(element interface{}) {
	q.array = append(q.array, element)
}

// Dequeue returns and removes an element from the front of the queue
// or error if the queue is empty.
func (q *Queue) Dequeue() (element interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	element = q.array[0]
	q.array[0] = nil // set to nil to free memory for element
	q.array = q.array[1:]
	return element, nil
}

// Peek returns, but does not remove, an element from the front of the
// queue or error if the queue is empty.
func (q *Queue) Peek() (element interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	element = q.array[0]
	return element, nil
}

// String returns the queue as a string.
func (q Queue) String() string {
	return fmt.Sprintf("size = %d values = %v", q.Size(), q.array)
}
