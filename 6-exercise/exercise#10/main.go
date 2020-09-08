package main

import "fmt"

func main() {
	counter := ticker()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func ticker() func() int {
	ticker := 0
	adder := func() int {
		ticker++
		return ticker
	}
	return adder
}
