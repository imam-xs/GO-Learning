package main

import "fmt"

func add(number1 int, number2 int) {
	sum := number1 + number2
	fmt.Println("Sum:", sum)
}

func main() {
	a := 10
	b := 20

	add(a, b)

	add(30, 40)
}