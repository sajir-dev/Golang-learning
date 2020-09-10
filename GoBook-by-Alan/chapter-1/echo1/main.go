package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	// for i := 1; i < len(os.Args); i++ {
	// 	s += os.Args[i]
	// }
	fmt.Println(s)
	fmt.Println(len(os.Args))
}

// strings.Join()
// fmt.Println(strings.Join(os.Args[1:], " "))
