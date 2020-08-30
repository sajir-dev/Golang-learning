package main

import "fmt"

func main() {
	student := struct {
		name  string
		grade string
	}{
		name:  "Raju",
		grade: "6",
	}

	fmt.Println(student.name, student.grade)
}
