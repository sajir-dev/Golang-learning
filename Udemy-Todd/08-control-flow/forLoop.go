package main

import "fmt"

func main() {

	i := 0

	for {
		if i == 3 {
			break
		}
		fmt.Println(i, "once")
		i++
	}

	// fmt.Println("Hello")

}
