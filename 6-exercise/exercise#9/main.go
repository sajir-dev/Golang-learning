package main

import "fmt"

func main() {

	fmt.Println(squareMinusNum(square, 8))

}

func square (int num) int {
	return num*num
}

func squareMinusNum (f(int num) int, num int) int{
	x:= f(num)-num
	return x
}