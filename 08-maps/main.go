package main

import "fmt"

func main() {
	// define map
	// emails := make(map[string]string)

	// Assing kv
	// emails["Bob"] = "bob@email.com"
	// emails["sherin"] = "sharon@email.com"
	// emails["Ravi"] = "rav@email.com"

	emails := map[string]string{"Bob": "bob@email.com", "sharon": "sharon@email.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// delete one value
	delete(emails, "Bob")
	fmt.Println(emails)
}

