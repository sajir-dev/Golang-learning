package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	// answer - 1
	// go func() {
	// 	c <- 42
	// }()

	// answer-2
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}
