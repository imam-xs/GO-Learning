package main

import "fmt"

// if a function has a name and is not associated with any type, it is called a standard function.
func standardFunction(message string) {
	fmt.Println(message)
}

func main() {
	standardFunction("This is an example of a standard function.")
}
