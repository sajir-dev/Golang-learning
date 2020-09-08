package main

import "fmt"

func multiplesOfThreeAndFive(num int) [2]int {
	sum := 0
	prod := 1
	for i := 3; i <= num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
			prod *= i
		}
	}
	var arr [2]int
	arr[0] = sum
	arr[1] = prod
	return arr
}

func main() {
	fmt.Println(multiplesOfThreeAndFive(10))
}
