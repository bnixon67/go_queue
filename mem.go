package main

import (
	"fmt"
	queue "github.com/bnixon67/go_queue/list_queue"
	"runtime"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats

	runtime.GC()

	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tHeapObjects = %v", m.HeapObjects)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	PrintMemUsage()

	fmt.Println("*** create queue")
	q := queue.New()
	PrintMemUsage()

	fmt.Println("*** enqueue 1000000 int")
	for n := 0; n < 1000000; n++ {
		q.Enqueue(n)
	}
	PrintMemUsage()

	fmt.Println("*** dequeue 1000000 int")
	for n := 0; n < 1000000; n++ {
		q.Dequeue()
	}
	runtime.GC()
	PrintMemUsage()
}
