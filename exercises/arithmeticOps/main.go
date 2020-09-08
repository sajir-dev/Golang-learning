package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func subtract(num1 int, num2 int) int {
	return num1 - num2
}

func multiply(num1 int, num2 int) int {
	return num1 * num2
}

func divide(num1 float32, num2 float32) float32 { // Integer division will not yeild decimal values in Go
	return num1 / num2
}

func modules(num1 int, num2 int) int {
	return num1 % num2
}

func greaterThanTest(num1 int, num2 int) bool {
	return num1 > num2
}

func lessThanTest(num1 int, num2 int) bool {
	return num1 < num2
}

func equlaityTest(num1 int, num2 int) bool {
	return num1 == num2
}

func main() {
	fmt.Printf("Sum: %d \n", add(3, 8))
	fmt.Printf("Diff: %d \n", subtract(4, 10))
	fmt.Printf("Product: %d \n", multiply(2, 24))
	fmt.Printf("Quotient: %f \n", divide(1, 15))
	fmt.Printf("Modules: %d \n", modules(20, 6))
	fmt.Printf("Greater than: %v\n", greaterThanTest(10, 10))
	fmt.Printf("Less than test: %v\n", lessThanTest(2, 5))
	fmt.Printf("Equality test: %v\n", equlaityTest(5, 5))
}
