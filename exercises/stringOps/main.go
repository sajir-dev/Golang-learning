//concatenate / reverse / reverse concat / mix the letters / split in half / remove all white spaces / add a character after every specific letter

package main

import (
	"fmt"
)

// concatentate
func concat(string1 string, string2 string) string {
	return string1 + string2
}

// Reverse
func reverse(word string) string {
	// fmt.Println(word)
	result := ""
	// fmt.Printf("%T, %T\n", word, result)
	for _, v := range word {
		result = string(v) + result // each step is added to the left side of the variable ie; a, ba, cba, dcba, edcba
		// fmt.Println(result)
	}
	return result
}

// func reverse(string word) string {
// 	var rev string
// 	for i := 0; i < len(word); i++ {
// 		rev[i] = word[len(word)-1-i]
// 	}
// }

// reverse concatenate
func revconcat(string1 string, string2 string) string {
	concatenatedStr := concat(string1, string2)
	rev := reverse(concatenatedStr)
	return rev
}

// mix the letters
func mixTheLetters(string1 string, string2 string) string {
	r1 := []rune(string1)
	r2 := []rune(string2)
	string3 := concat(string1, string2)
	mix := []rune(string3)
	fmt.Println(mix)
	for i := 0; i < len(mix); i++ {
		if i%2 == 0 {
			mix[i] = r1[(i / 2)]
			fmt.Println(i, mix[i])
		} else {
			mix[i] = r2[(i / 2)]
			fmt.Println(i, mix[i])
		}
	}
	return string(mix)
}

// split in half
func splitHalf(word string) (string, string) {
	r := []rune(word)
	r1 := r[0 : len(r)/2]
	r2 := r[len(r)/2 : len(r)]
	return string(r1), string(r2)
}

//remove whitespaces
func removeSpace(word string) string {
	r := []rune(word)
	for i := 0; i < len(r); i++ {
		if r[i] == 32 {
			// r[i] = r[len(r)-1]
			// r = r[:len(r)-1]
			// r[len(r)-1] = r[i]
			r1 := r[:i]
			r2 := r[i+1 : len(r)]
			r = append(r1, r2...) // adding two runes together
			i--
		}
	}
	return string(r)
}

// add a char after every specific letter
func addChar(str string, ch string, chAdd string) string {
	rStr := []rune(str)
	rChar := []rune(ch)
	rChAdd := []rune(chAdd)
	for i := 0; i < len(rStr); i++ {
		if rStr[i] == rChar[0] {
			println("here")
			r1 := rStr[:i]
			// r2 := rChAdd[0]
			r2 := rStr[i:len(rStr)]

			rStr = append(r1, rChAdd[0])
			fmt.Println(i, rStr)
			rStr = append(rStr, r2...)
			i--
		}
	}
	return string(rStr)
}

func main() {
	string1 := "abc"
	string2 := "bar"
	r := []rune(string1)

	// var string3 string = concat(string1, string2)
	string3 := concat(string1, string2)
	// fmt.Println(len(string3))
	fmt.Println(string3)
	fmt.Println(reverse("abcde"))
	fmt.Println(revconcat("!dlrow ", "olleh"))
	fmt.Println(string(r))
	fmt.Println(mixTheLetters("abcd", "pqrs"))
	fmt.Println(splitHalf("abcd"))
	space := " "
	r2 := []rune(space)
	fmt.Println(r2)
	// r[1] = r[2]
	// r = r[:len(r)-1]
	// fmt.Println(len(r))
	// fmt.Println(string(r))
	fmt.Println(removeSpace("abc   def"))
	fmt.Println(addChar("pekaboo", "o", "b"))
}
