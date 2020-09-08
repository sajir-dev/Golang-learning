package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	str := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(str)
	fmt.Printf("%T\n", str)
	fmt.Printf("%T\n", bs)

	// people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for _, v := range people {
		fmt.Println("First: ", v.First)
		fmt.Println("Last: ", v.Last)
		fmt.Println("Age: ", v.Age)
		fmt.Println("----------")
	}

}
