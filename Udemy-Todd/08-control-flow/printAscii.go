package main

import "fmt"

func main() {
	i := 33

	for {
		fmt.Printf("%v : %#U\n", i, i)
		i++
		if i > 122 {
			break
		}
	}
}
