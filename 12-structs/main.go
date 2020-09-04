package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// value reciever method
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I'm " + strconv.Itoa(person.age)
}

// pointer reciever
func (person *Person) hasBirthday() {
	person.age++
}

// get married
func (person *Person) getMarried(husband string) {
	person.lastName = husband
}

func main() {
	// init a person using string
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 30}

	// another way of declaration
	// person1 := Person{"samantha", "smith", "boston", "f", 34}

	fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("John")
	fmt.Println(person1.greet())
}

