package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++ // func text converts to string
		// fmt.Println(counts)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Printf conversions (verbs)
//  %d 				decimal integer
//  %x, %o, %b		integer in hexadecimal
//  %f, %g, %e		floating point number
// 					3.141593, 3.14159265258973, 3.141593e+00
//  %t 				boolean
//  %c				rune (Unicode point)
//  %s				string
//  %q				quoted string
//  %v				any value in a natural format
//  %T				type of any value
//  %%				literal percent sign
