package main

import (
	"fmt"
	"math"
)

//Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	// circle := Circle{x: 0, y: 0, radius: 5}
	circle := Circle{0, 0, 5}
	// rectangle := Rectangle{width: 5, height: 10}
	rectangle := Rectangle{4, 10}

	fmt.Println("Area of circle: ", getArea(circle))
	fmt.Println("Area of rectangle: ", getArea(rectangle))
}
