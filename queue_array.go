package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Queue struct {
	array []interface{}
}

func NewQueue() Queue {
	return Queue{array: make([]interface{}, 0)}
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

func main() {
	fmt.Println("*** create queue")
	q := NewQueue()
	fmt.Println("q:", q, "size:", q.Size())
	fmt.Println()

	fmt.Println("*** dequeue from empty queue")
	element, err := q.Dequeue()
	fmt.Println("element:", element, "type:", reflect.TypeOf(element), "err:", err)
	fmt.Println()

	fmt.Println("*** enqueue 1")
	q.Enqueue(1)
	fmt.Println("q:", q, "size:", q.Size())
	fmt.Println()

	fmt.Println("*** enqueue \"hello\"")
	q.Enqueue("hello")
	fmt.Println("q:", q, "size:", q.Size())
	fmt.Println()

	fmt.Println("*** dequeue")
	element, err = q.Dequeue()
	fmt.Println("element:", element, "type:", reflect.TypeOf(element), "err:", err)
	fmt.Println("q:", q, "size:", q.Size())
	fmt.Println()

	fmt.Println("*** peek")
	element, err = q.Peek()
	fmt.Println("element:", element, "type:", reflect.TypeOf(element), "err:", err)
	fmt.Println("q:", q, "size:", q.Size())
	fmt.Println()

	fmt.Println("*** dequeue")
	element, err = q.Dequeue()
	fmt.Println("element:", element, "type:", reflect.TypeOf(element), "err:", err)
	fmt.Println("q:", q, "size:", q.Size())
	fmt.Println()
}
