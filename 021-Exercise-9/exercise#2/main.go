package main

import "fmt"

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hi there, my name is ", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Aju"}
	// saySomething(p) - will not work because the reciever is a pointer type
	saySomething(&p1)
	p1.speak()
}
