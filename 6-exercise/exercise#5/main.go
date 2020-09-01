package main

import (
	"fmt"
	"math"
)

// Circle is a struct
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Rectangle is a struct
type Rectangle struct {
	length, width float64
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

// Shape is an interface
type Shape interface {
	area() float64
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c := Circle{
		radius: 21.3,
	}

	r := Rectangle{
		length: 12,
		width:  8,
	}

	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}
