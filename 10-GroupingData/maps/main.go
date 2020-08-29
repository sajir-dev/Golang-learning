package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Bond":  28,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	// Checking the value of the key
	v, ok := m["Barnes"]

	// if v, ok := m["Barnes"]; ok {
	// 	fmt.Println("The if print", v)
	// }

	// fmt.Println(m)
	fmt.Println(ok, v)

	// adding a value to map
	m["eka"] = 33

	// printing map using for k, v
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Deleting from map
	// delete(m, "James")
	// fmt.Println(m)

	// We can also delete an entry that doesnt exist
	// We use ok idiom to avoid that
	if v, ok := m["Bond"]; ok {
		fmt.Println("value: ", v)
		delete(m, "Bond")
	}

	fmt.Println(m)

	// slice using template literal
	xi := []int{2, 43, 3, 4, 5, 13, 44, 32}

	// Printing a slice using for range
	for k, v := range xi {
		fmt.Println(k, v)
	}

	// printing a slice values only usng for range
	for _, v := range xi {
		fmt.Println(v)
	}

	// printing a slice key only using for range
	for k := range xi {
		fmt.Println(k)
	}
}
