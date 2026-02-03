package main

import "fmt"

func welcomeMessage() {
	fmt.Println("Welcome to the application!")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)
	return num1, num2
}

func addNumbers(a int, b int) int {
	return a + b
}

func displayResult(name string, sum int) {
	fmt.Println("hello", name)
	fmt.Println("The sum of the two numbers is:", sum)
}

func goodbyeMessage() {
	fmt.Println("Thank you for using the application.")
	fmt.Println("Goodbye!")
}

func main() {
	welcomeMessage()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := addNumbers(num1, num2)
	displayResult(name, sum)
	goodbyeMessage()
}
