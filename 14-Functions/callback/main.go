package main

import "fmt"

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	sumOfAll := sum(arr...)

	fmt.Println(sumOfAll)
	fmt.Println("Sum of all: ", sum(arr...))
	fmt.Println("Sum of Even: ", evenSum(sum, arr...))
	fmt.Println("Some of Odd: ", oddSum(sum, arr...))
}

func sum(arrSlice ...int) int {
	total := 0
	for _, v := range arrSlice {
		total += v
	}
	return total
}

func evenSum(someSumFn func(someSlice ...int) int, anotherSlice ...int) int {
	var evenSlice []int
	for _, v := range anotherSlice {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	sum := someSumFn(evenSlice...)
	return sum
}

func oddSum(someSumFn func(someSlice ...int) int, anotherSlice ...int) int {
	var oddSlice []int
	for _, v := range anotherSlice {
		if v%2 != 0 {
			oddSlice = append(oddSlice, v)
		}
	}

	return someSumFn(oddSlice...)
}
