package main

import "fmt"

func primeChecker(num int) bool {
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeChecker(23))
}
