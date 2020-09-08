package main

import "fmt"

func main() {
	// My solution
	// student := struct {
	// 	name  string
	// 	grade string
	// }{
	// 	name:  "Raju",
	// 	grade: "6",
	// }

	// fmt.Println(student.name, student.grade)

	// Todd's solution
	s := struct {
		name      string
		friends   map[string]int
		favDrinks []string
	}{
		name: "James Bonda",
		friends: map[string]int{
			"moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"martini",
			"Water",
		},
	}

	fmt.Println(s.name)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}
}
