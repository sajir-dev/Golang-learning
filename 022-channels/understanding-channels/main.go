package main

import "fmt"

func main() {
	// c := make(chan int)

	// // c <- 42
	// // will not work
	// // because fatal error: all goroutines are asleep - deadlock!

	// // will work
	// go func() {
	// 	c <- 42
	// }()

	// fmt.Println(<-c)

	c := make(chan int, 2)

	c <- 42
	c <- 43 // will not work because of deadlock

	fmt.Println(<-c)
	fmt.Println(<-c)
}
