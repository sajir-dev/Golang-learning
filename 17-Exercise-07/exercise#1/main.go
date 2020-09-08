package main

import "fmt"

// Person type
type Person struct {
	first string
	last  string
}

func (person Person) changeMeByValue(name string) {
	person.first = name
	fmt.Println("Inside value reciever fn", person)
}

func (person *Person) changeMeByReference(name string) {
	person.first = name
	fmt.Println("Inside pointer reciever fn", person)
}

func main() {
	p1 := Person{
		"Taylor", "Swift",
	}

	fmt.Println(p1)

	p1.changeMeByValue("Brooke")

	fmt.Println(p1)

	p1.changeMeByReference("Ashlin")

	fmt.Println(p1)
}
