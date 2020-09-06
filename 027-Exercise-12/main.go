package main

import (
	"fmt"

	"./dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(4),
	}
	fmt.Println(fido)
}
