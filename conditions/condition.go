package main

import "fmt"

func main(){
	//if-else condition
	age := 18

	if age > 18 {
		fmt.Println("You are an adult person.")
	} else if age <= 18 {
		fmt.Println("You are not an adult person.")
	} else {
		fmt.Println("Invalid age.")
	}

	//use && and || operators
	nationality := "Bangladeshi"
	languageSkill := "Golang"
	
	if nationality == "Bangladeshi" && languageSkill == "Golang" {
		fmt.Println("You are eligible for the job.")
	} else {
		fmt.Println("You are not eligible for the job.")
	}

	if nationality == "Bangladeshi" || languageSkill == "Python" {
		fmt.Println("You have some qualifications.")
	} else {
		fmt.Println("You do not have the required qualifications.")
	}
}