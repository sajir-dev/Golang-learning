package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Sayings []string
}

func main() {
	p1 := person {
		First:	"James",
		Last:	"Bond",
		Sayings:	[]string{
			"Shaken, not stirred",
			"Any last wishes?",
			"Never say never",
		}
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		// fmt.Println(err)
		// return
		log.Fatalln("JSON did not marshal - here's the error: ", err)
	}

	fmt.Println("string(bs)")
}