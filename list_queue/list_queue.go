// Package list_queue is an implementation of a queue data structure in go backed by an list
package list_queue

import (
	"errors"
)

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	first *Node
	last  *Node
	size  int
}

func NewQueue() Queue {
	return Queue{first: nil, last: nil, size: 0}
}

func (q *Queue) IsEmpty() bool {
	return q.Size() <= 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Enqueue(element interface{}) {
	node := Node{value: element, next: nil}

	if q.first == nil {
		q.first = &node
		q.last = &node
	} else {
		q.last.next = &node
		q.last = &node
	}
	q.size++
}

func (q *Queue) Dequeue() (element interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	first := q.first

	element = first.value

	q.first = q.first.next

	first = nil
	q.size--

	return element, nil
}

func (q *Queue) Peek() (element interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	element = q.first.value
	return element, nil
}
