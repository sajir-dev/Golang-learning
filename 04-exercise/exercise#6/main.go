package main

import "fmt"

func main() {
	names := []string{"James", "Bond", "Shaken not stirred"}
	class := []string{"Miss", "MoneyPenny", "Hellooo", "James"}
	ss := [][]string{names, class}
	for i := range ss {
		for j, u := range ss[i] {
			fmt.Println(i, j, u)
		}
	}
}
