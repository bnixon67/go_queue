// Package array_queue is an implementation of a queue data structure in go backed by an array
package array_queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	array []interface{}
}

func NewQueue() *Queue {
	return &Queue{array: make([]interface{}, 0)}
}

func (q *Queue) IsEmpty() bool {
	return q.Size() <= 0
}

func (q *Queue) Size() int {
	return len(q.array)
}

func (q *Queue) Enqueue(element interface{}) {
	q.array = append(q.array, element)
}

func (q *Queue) Dequeue() (element interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	element = q.array[0]
	q.array[0] = nil // set to nil to free memory for element
	q.array = q.array[1:]
	return element, nil
}

func (q *Queue) Peek() (element interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	element = q.array[0]
	return element, nil
}

func (q Queue) String() string {
	return fmt.Sprintf("size = %d values = %v", q.Size(), q.array)
}
