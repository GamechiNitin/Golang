package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")

	logInCount := 23
	var message string

	if logInCount < 10 {
		message = "Regular user"
	} else if logInCount > 10 {
		message = "Watch out"
	} else {
		message = "Exactly 10 LogIn count"
	}

	fmt.Println(message)

	//:::::::::::::: Checking condition

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	//:::::::::::::: Initilize and Checking condition

	if num := 3; num < 10 {
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is not less then 10")
	}

}
