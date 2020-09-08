package main

import "fmt"

type customErr struct {
	info string
}

// implementing the Error interface
func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Println("foo ran", e, "\n", e.(customErr).info)
}

func main() {
	err1 := customErr{
		info: "Need more coffee",
	}

	foo(err1)
}
