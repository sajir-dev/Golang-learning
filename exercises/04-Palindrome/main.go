package main

import "fmt"

func palindromeCheck(str string) bool {
	result := ""
	for _, v := range str {
		result = string(v) + result
	}
	if result == str {
		return true
	}
	return false
}

func main() {
	fmt.Println(palindromeCheck("abdgddasdfqewba"))
}
