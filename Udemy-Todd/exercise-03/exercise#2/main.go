package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
