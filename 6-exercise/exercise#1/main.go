package main

import "fmt"

func main() {
	n := foo(2)
	x, y := bar(2, "undampori")
	fmt.Println(n)
	fmt.Println(x, y)
}

func foo(a int) int {
	return 2 * a
}

func bar(a int, s string) (string, int) {
	return s, 2 * a
}
