package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "Aby",
		lastName:  "Jon",
		iceCream:  []string{"ButterSctoch", "Strawberry", "Chocolate"},
	}

	p2 := person{
		firstName: "Aju",
		lastName:  "Jon",
		iceCream:  []string{"Vanilla", "Dry fruits", "Mango"},
	}

	fmt.Printf("%v %v \n", p1.firstName, p1.lastName)

	for _, v := range p1.iceCream {
		fmt.Println(v)
	}

	fmt.Printf("%v %v \n", p2.firstName, p2.lastName)

	for _, v := range p2.iceCream {
		fmt.Println(v)
	}
}
