package main

import "fmt"

const (
	y0 = 2017 + iota
	y1 = 2017 + iota
	y2 = 2017 + iota
)

func main() {
	fmt.Println(y0)
	fmt.Println(y1)
	fmt.Println(y2)
}
