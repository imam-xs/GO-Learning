package main
import "fmt"

num := 10

func main() {
	age := 25
	if age > 20 {
		num := 15
		fmt.Println("Inner num:", num) // Outputs: Inner num: 15
	}
	fmt.Println("Outer num:", num) // Outputs: Outer num: 10
}