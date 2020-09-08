package main

import "fmt"

func main() {
	m := map[string][]string{
		`Rudy`:  []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		`Judy`:  []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		`no_dr`: []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}

	// for k := range m {
	// 	fmt.Println(k, m[k])
	// }

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
