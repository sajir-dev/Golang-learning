package main

import "fmt"

func main() {
	defer bar()

	foo()

}

func foo() {
	fmt.Println("Hi there")
}

func bar() {
	fmt.Println("Hi, I'm the defer")
}
