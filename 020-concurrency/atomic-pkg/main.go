package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// go run -race main.go can give us info if it has a race condition

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter: \t", atomic.LoadInt64(&counter))
			wg.Done()

		}()
		fmt.Println("Goroutines:\t ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
