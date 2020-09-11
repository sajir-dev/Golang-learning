package main

import "fmt"

func main() {
	a := 12
	for i := 2; i < a/2; i++ {
		if a%i == 0 {
			fmt.Println("Number is not prime")
			return
		}
	}
	fmt.Println("Number is prime")
}
