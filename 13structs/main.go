package main

import "fmt"

func main() {
	fmt.Println("Struct in Golang")

	// No inheritance, super or parent or child in golang

	nitinGamechi := User{"Nitin Gamechi", "nitingamechi@flutter.dev", true, 19}
	fmt.Println(nitinGamechi)

	// Output formating
	fmt.Println("\n:::::::::::::: Output formating ::::::::::::::")
	fmt.Printf("Nitin details are : %+v\n", nitinGamechi)
	fmt.Printf("Name is %v, and email is %v\n", nitinGamechi.Name, nitinGamechi.Email)

}

// Define Struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
