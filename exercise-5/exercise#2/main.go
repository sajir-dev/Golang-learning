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
		lastName:  "Tony",
		iceCream: []string{
			"ButterSctoch",
			"Strawberry",
			"Chocolate",
		},
	}

	p2 := person{
		firstName: "Aju",
		lastName:  "Jon",
		iceCream: []string{
			"Vanilla",
			"Dry fruits",
			"Mango",
		},
	}

	// my solution
	// m := map[string]person{"Tony": p1, "Jon": p2}
	// fmt.Println(m)

	// real solution
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range m {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.iceCream {
			fmt.Println(i, val)
		}
		fmt.Println("------")
	}

}
