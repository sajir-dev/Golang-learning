package main

import "fmt"

func main() {
	slice := []int{
		2, 3, 5, 5, 39, 20,
	}

	s := add(slice...)
	fmt.Println(s)

	p := add(slice)
	fmt.Println(p)
}

func add(slice ...int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
