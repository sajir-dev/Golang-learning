package main

import "fmt"

func arrSum(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	num := [5]int{2, 3, 5, 0, 1}
	fmt.Println(arrSum(num))
}
