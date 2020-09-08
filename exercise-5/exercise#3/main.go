package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: false,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "grey",
		},
		luxury: true,
	}

	fmt.Println(t1)
	fmt.Println(s1)

	fmt.Println(t1.doors)
	fmt.Println(s1.color)
}
