package main

import "fmt"

// constants can be declared using const keywords and type gets assigned automatically
const (
	a = "abc"
	b = 24
	c = false
)

func main() {
	fmt.Printf("a value: %v \t type: %T", a, a)
	fmt.Printf("b value: %v \t type: %T", b, b)
	fmt.Printf("c value: %v \t type: %T", c, c)
}
