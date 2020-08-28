package main

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	println(greeting("Sajir"))
	println(getSum(3, 4))
}
