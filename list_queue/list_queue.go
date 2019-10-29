// Package list_queue implements a queue backed by an list.
// A queue is a First In-First Out (FIFO) data type, i.e. the element enqueued
// first will be dequeued first.
package list_queue

import (
	"errors"
	"fmt"
)

// node represents an item in the queue.
type node struct {
	// Value is the queue item.
	value interface{}
	// Next is a pointer to the next item in the queue
	next *node
}

// Queue represents a queue backed by a list.
type Queue struct {
	// First and last points to the first and last node in the queue.
	first, last *node
	// Size contains the number of items in the queue.
	size int
}

// New returns an initialized empty queue.
func New() *Queue {
	return &Queue{first: nil, last: nil, size: 0}

}

// IsEmpty reports whether the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.Size() <= 0
}

// Size reports the number of elements are in the queue.
func (q *Queue) Size() int {
	return q.size
}

// Enqueue adds the element to the end of the queue.
func (q *Queue) Enqueue(element interface{}) {
	// create the node
	node := node{value: element, next: nil}

	if q.first == nil {
		// if queue is empty, then first and last is the same node
		q.first = &node
		q.last = &node
	} else {
		// add node to end of queue
		q.last.next = &node
		// update last to point to new end of queue
		q.last = &node
	}

	// update size of queue
	q.size++
}

// Dequeue returns and removes an element from the front of the queue
// or error if the queue is empty.
func (q *Queue) Dequeue() (element interface{}, err error) {
	// return error is queue is empty
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	// get first node of the queue
	first := q.first

	// get element from node
	element = first.value

	// advance the queue to the next element
	q.first = q.first.next

	// free to node to prevent memopry leak
	first = nil

	// update size of queue
	q.size--

	return element, nil
}

// Peek returns, but does not remove, an element from the front of the
// queue or error if the queue is empty.
func (q *Queue) Peek() (element interface{}, err error) {
	// return error is queue is empty
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	// get element from head of the queue
	element = q.first.value

	return element, nil
}

// String returns the queue as a string.
func (q Queue) String() string {
	// values will store the queue as a list
	var values []interface{}

	// iterate thru the queue
	for node := q.first; node != nil; node = node.next {
		values = append(values, node.value)
	}

	return fmt.Sprintf("size = %d values = %v", q.size, values)
}
