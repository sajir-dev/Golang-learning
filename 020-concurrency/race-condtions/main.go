package main

import (
	"fmt"
	"runtime"
	"sync"
)

// TO PREVENT THE DATA RACE WE SHOULD ADD MUTEX
// go run -race main.go can give us info if it has a race condition

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()

		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
