package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is ", p.first, " and my age is ", p.age)
}

func main() {
	p1 := person{
		first: "Salim",
		last:  "Kumar",
		age:   35,
	}

	p1.speak()
}
