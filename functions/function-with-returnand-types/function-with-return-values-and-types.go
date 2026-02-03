package main

import "fmt"

// Function that adds two integers and returns the sum as an integer that is why the return type is specified as int
func addInts (num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

//example of function with multiple return types
func getNumbers(number1 int, number2 int) (int, int) {
	sumation := number1 + number2
	multification := number1 * number2
	return sumation, multification
}

//more examples
func sayHello(name string) {
	fmt.Println("Hello,", name)
}

func welcomeMessage() {
	fmt.Println("Welcome to Go programming!")
}

func main () {
	a := 10
	b := 20

	result := addInts(a, b)
	fmt.Println("Sum:", result)

	// Calling function with multiple return types
	sumation, multification := getNumbers(5, 5)
	fmt.Println("Sumation:", sumation)
	fmt.Println("Multification:", multification)

	// Calling sayHello function
	sayHello("Emon")

	// Calling welcomeMessage function
	welcomeMessage()
}