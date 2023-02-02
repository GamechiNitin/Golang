package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	
	// var ptr *int
	// fmt.Println("Value of pointers", ptr)

	myNumber := 24

	var ptr = &myNumber
	fmt.Println("Memory address of pointers : ", ptr) // Memory address
	fmt.Printf("Type of pointers : %T \n", ptr)       // Type address
	fmt.Println("Value of pointers : ", *ptr)         // Value

	*ptr = *ptr + 26
	fmt.Println("New Value of pointers : ", *ptr)     // Value
	fmt.Println("New Value of variable : ", myNumber) // Value

}
