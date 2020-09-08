package main

import "fmt"

func main() {
	f := playGuitar()
	fmt.Println(f())

	g := str1PlayGuitar("Hawk")
	fmt.Println(g())

	h := str2PlayGuitar()
	fmt.Println(h("Rock"))
}

// returing without args
func playGuitar() func() string {
	return func() string {
		return " playing lalalala"
	}
}

// outerfn with args
func str1PlayGuitar(str string) func() string {
	return func() string {
		return str + " playing lalalala"
	}
}

// innerfn with args
func str2PlayGuitar() func(str string) string {
	return func(str string) string {
		return str + " playing lalalala"
	}
}
