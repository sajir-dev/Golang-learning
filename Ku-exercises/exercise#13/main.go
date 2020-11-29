package main

import "fmt"

func main() {
	limit := 30
	primesSlice := []int{}

	for j := 2; j <= limit; j++ {
		if checkPrime(j) {
			primesSlice = append(primesSlice, j)
		}
	}
	fmt.Println("Prime numbers upto", limit, primesSlice)
}

func checkPrime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
