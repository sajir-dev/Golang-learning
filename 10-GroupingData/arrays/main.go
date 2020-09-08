package main

import "fmt"

func main() {
	var arr [5]int
	arr[3] = 32
	fmt.Println(arr)
	fmt.Println(len(arr))
}
