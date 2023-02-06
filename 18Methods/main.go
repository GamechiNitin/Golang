package main

import "fmt"

func main() {
	fmt.Println("Struct in Golang")

	// No inheritance, super or parent or child in golang

	nitinGamechi := User{"Nitin Gamechi", "nitingamechi@flutter.dev", true, 19}
	fmt.Println(nitinGamechi)

	// Output formating
	fmt.Println("\n:::::::::::::: Struct Output formating ::::::::::::::")
	fmt.Printf("Nitin details are : %+v\n", nitinGamechi)

	fmt.Println("\n:::::::::::::: Methods Output formating ::::::::::::::")
	nitinGamechi.GetStatus()

	fmt.Println("\n:::::::::::::: Change Email Output formating ::::::::::::::")
	nitinGamechi.NewMail()
	fmt.Printf("Nitin details are : %+v\n", nitinGamechi) // This does not change original data

}

// Define Struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active :", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New Email is :", u.Email)
}
