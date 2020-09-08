package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("a: %v, b: %v \n", a, b)

	fmt.Printf("typeof- b: %T\n", b)

	fmt.Printf("%v \n", &a)

	fmt.Printf("%v \n", *&a)

	// change the value with pointer
	*b = 10
	fmt.Println("a =", a)
}
