package main  //every Go program belongs to a package. without package main, the program would not start.
import "fmt" //the fmt package is part of Go’s standard library and is used for formatting input and output

func main(){  //a Go program execution always starts from main() function
	fmt.Println("Hello World")
}