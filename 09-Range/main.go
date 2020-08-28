package main

import "fmt"

func main() {
	// Declaring a slice
	ids := []int{33, 76, 54, 23, 11, 2}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	// looping without index
	for _, id := range ids {
		fmt.Printf("ID: %d \n", id)
	}

	// Add all ids together
	var sum int // or we can use sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// range with maps
	emails := map[string]string{"Bob": "bob@wmail.com", "sharon": "sharon@email.com"}
	for k, v := range emails {
		fmt.Printf("%s : %s \n", k, v)
	}

	// range with only map keys
	// emails := map[string]string{"Bob": "bob@wmail.com", "sharon": "sharon@email.com"}
	for k := range emails {
		fmt.Printf("name: %s\n", k)
	}

	// range with only map values
	// emails := map[string]string{"Bob": "bob@wmail.com", "sharon": "sharon@email.com"}
	for _, v := range emails {
		fmt.Printf("email: %s\n", v)
	}
}
