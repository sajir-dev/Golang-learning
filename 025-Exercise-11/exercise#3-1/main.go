package main

import "fmt"

type hotdog int

func main() {
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	// y = x not possible because cannot use x (type hotdog) as type int in assignment
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
