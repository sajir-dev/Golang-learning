package main

import "fmt"

func multiplyByK(arr *[3]int, k int) {
	// var newArr[3] int
	for i := 0; i < 3; i++ {
		arr[i] = k * arr[i]
	}
	// return

	// arr[0] = k * arr[0]
}

func main() {
	arr := [3]int{4, 7, 10}
	multiplyByK(&arr, 2)
	fmt.Println(arr)
	// ptr := &arr
	// fmt.Println(multiplyByK(ptr, 2))
}
