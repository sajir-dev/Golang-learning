package main

import "fmt"

func main() {

	c := func(name string) string {
		// fmt.Sprint prints to a string
		return fmt.Sprint("Hello " + name)
	}("George")

	fmt.Println(c)
}
