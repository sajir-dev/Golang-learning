package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	// func CompareHashAndPassword(hashedPassword, password []byte) error
	pw := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(pw))
	if err != nil {
		fmt.Println(err)
		fmt.Println("WRONG PASSWORD")
		return
	}

	fmt.Println("YOU ARE LOGGED IN")

}
