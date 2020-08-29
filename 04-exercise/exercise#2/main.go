package main

import "fmt"

func main() {
	arr := []int{53, 32, 432, 24, 432, 5, 85, 78, 87, 78}
	for k, v := range arr {
		fmt.Println(k, v)
	}

	fmt.Printf("%T", arr)
}
