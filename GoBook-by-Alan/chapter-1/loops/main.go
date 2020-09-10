package main

import "fmt"

func main() {
	// while loop
	i := 2
	for i < 8 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	for {
		fmt.Println(i)
		i++
		if i == 30 {
			fmt.Println("breaking infinte loop")
			break
		}
	}
}

// different variable declarations
//  s := "" // can be used inside a function only
//  var s string
//  var s = ""
//  var s string = ""
