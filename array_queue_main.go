package main

import (
	"fmt"
	"github.com/bnixon67/go_queue/array_queue"
	"reflect"
)

func main() {
	fmt.Println("*** create queue")
	q := array_queue.NewQueue()
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
