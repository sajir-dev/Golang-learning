package main

import "fmt"

func main() {
	// a := 42
	// fmt.Println(a)
	// fmt.Println(&a) // & gives us the address

	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T \n", &a)

	// b := &a
	// fmt.Println(b)
	// fmt.Println(*b) // * gives us the value in the address
	// fmt.Println(*&a)

	// *b = 43
	// fmt.Println(a)

	// When to use pointers?
	x := 0
	fmt.Println(x)
	foo(&x)
	fmt.Println(x)
}

func foo(p *int) {
	*p = 1
}
