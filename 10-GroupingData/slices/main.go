package main

import "fmt"

func main() {
	// x:= type{ } composite literals
	x := []int{4, 5, 6, 7, 9}
	// fmt.Printf("%T, \t %v\n", x, x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))
	// fmt.Println(x[3])

	// append to a slice
	x = append(x, 12, 13, 14)
	fmt.Println(x)

	y := []int{100, 101, 102}
	x = append(x, y...)

	fmt.Println(x)
	// Deleting 3rd and 4th element
	x = append(x[:3], x[5:]...)

	fmt.Println(x)

	// using make fn
	z := make([]int, 10, 100)
	fmt.Println(y)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z[0] = 99
	z[9] = 999

	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z = append(z, x...)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
}
