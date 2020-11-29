package main

import "fmt"

func main() {
	for i := 2; i <= 12; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "X", j, "=", i*j)
		}
	}
}
