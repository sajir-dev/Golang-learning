package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("------")
	fmt.Printf("c: %T \n", c)
	fmt.Printf("cr: %T\n", cr)
	fmt.Printf("cs: %T\n", cs)

	// Assignment: c can be assigned to cr and cs because general can be assigned to specific channels. But at the same time cr, cs cannot be assigned to c, because it might lose some information like when we tried to assign float to an int.(direction info loses in channels, decimal value loses in float to int conversion)
	// Convert: general channel (here c) can be converted into directional, but opposite is not possible
}
