package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p1 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   42,
	}

	fmt.Println(sa1)
	fmt.Println(p1)

	// struct properties can be accessed with .
	// here the sa1.person.first is can be accessed with sa1.first because the first is getting promoted. it wasnt possible if the struct that creates sa1 has own a property named first
	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(p1.first, p1.last, p1.age)

}
