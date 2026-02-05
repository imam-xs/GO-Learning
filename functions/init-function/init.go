package main

import "fmt"

var globalVar = 50

func main() {
	fmt.Println(globalVar)
}

//init function: we cannot call the init function directly, it is automatically called by the Go runtime before the main function is executed. 
// The init function is used to initialize variables or perform setup tasks before the main function runs. 
func init() {
	fmt.Println(globalVar)
	globalVar = 100
}

//output:
// 50
// 100