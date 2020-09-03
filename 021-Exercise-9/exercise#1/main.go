package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	goOne()
	wg.Add(1)
	go goTwo()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	wg.Add(2)
	go goThree()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	wg.Wait()
}

func goOne() {
	for i := 0; i < 100; i++ {
		fmt.Println("go 1 loop: ", i)
	}
}

func goTwo() {
	for i := 0; i < 100; i++ {
		fmt.Println("go 2 loop: ", i)
	}
	wg.Done()
}

func goThree() {
	for i := 0; i < 100; i++ {
		fmt.Println("go 3 loop: ", i)
	}
	wg.Done()
}
