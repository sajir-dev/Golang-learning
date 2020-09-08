package main

import "fmt"

func main() {

	c := func(str string) string {
		return fmt.Sprint("You are so " + str)
	}

	// fmt.Println(c)
	fmt.Println(c("Cool"))
	fmt.Printf("%T: ", c)
}
