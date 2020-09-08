package main

import "fmt"

// var a int = 10

func main() {
	a := 5
	fmt.Printf("Value of a: \n demimal %d \n binary: %b\n hex: %#x\n", a, a, a)
	a = a << 3
	fmt.Printf("a: %v", a)
}
