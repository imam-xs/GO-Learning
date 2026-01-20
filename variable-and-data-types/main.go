package main

import "fmt"

func main()  {
	//explicit declaration (var)
	var age int = 30 //integer
	fmt.Println(age)

	var personName string = "Jon doe" // string
	fmt.Println(personName)

	var name = "Emon" //type inferred as string
	fmt.Println(name)

	var f float64 = 30.25 //floating number
	fmt.Println(f)

	var bol bool = true //boolean
	fmt.Println(bol)

   //short declaration (:=)
    clas := 10
	fmt.Println(clas)

	subject := "English"
	fmt.Println(subject)
	
}
