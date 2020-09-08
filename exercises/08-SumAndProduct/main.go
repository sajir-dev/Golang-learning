package main

import "fmt"

func sumOrProduct(num int) [2]int {
	sum := 0
	prod := 1
	for i := 1; i <= num; i++ {
		sum += i
		prod *= i
	}
	var arr [2]int
	arr[0] = sum
	arr[1] = prod
	return arr
}

func main() {
	fmt.Println(sumOrProduct(4)[0])
	fmt.Println(sumOrProduct(4)[1])
}
