package main

import "fmt"

func main() {
	m := map[string][]string{
		`Rudy`:  []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		`Judy`:  []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		`no_dr`: []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`dear jon`] = []string{`caterer`, `crushed`}

	for k, v := range m {
		fmt.Println("doc of: ", k)
		for i, u := range v {
			fmt.Println("\t", i, u)
		}
	}
}
