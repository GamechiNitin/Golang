package main

import "fmt"

const LogInToken string = "dcccnnc464s7s4c7a5s74dc7sdcnnvvdccddchvhxzxdcgdcsdj"

func main() {

	// :::::::::::::::::::::::::::::: Sting :::::::::::::::::::::::::::::::::::::::

	var userName string = "Nitin Gamechi"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	// :::::::::::::::::::::::::::::: Bool :::::::::::::::::::::::::::::::::::::::

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// :::::::::::::::::::::::::::::: Int :::::::::::::::::::::::::::::::::::::::

	//  Type uint8 Example : Range 0 - 255 (2^8)
	var value uint8 = 255
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T \n", value)

	//  Type uint16 Example : Range 0 - 65,536 (2^16)
	//  Type uint, int is most common used

	var value2 uint16 = 256
	fmt.Println(value2)
	fmt.Printf("Variable is of type: %T \n", value2)

	// :::::::::::::::::::::::::::::: Float :::::::::::::::::::::::::::::::::::::::

	//  Type float32 Example : 5 digit Precission after Decimal
	var smallFloat float32 = 255.20254545220
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//  Type float64 Example : 11 digit Precission after Decimal
	var largeFloat float64 = 255.20254545220
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)

	// ::::::::::::::::::::::::: Default Values and Alises ::::::::::::::::::::::::::::::::::::::::::::

	//  Type int no Initialize Example : Default value 0
	var anothevarible int
	fmt.Println(anothevarible)
	fmt.Printf("Variable is of type: %T \n", anothevarible)

	//  Type string no Initialize Example : Default value Empty
	var anothevarible2 string
	fmt.Println(anothevarible2)
	fmt.Printf("Variable is of type: %T \n", anothevarible2)

	// :::::::::::::::::::::::::::::: Implicit Type :::::::::::::::::::::::::::::::::::::::

	var website = "nitingamechi.in"
	fmt.Println(website)

	// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	// Here var can be initialize without var keyword. It cannot used in Gobal Variable
	tense := "tense.co"
	fmt.Println(tense)

	// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	fmt.Println(LogInToken)
	fmt.Printf("Variable is of type: %T \n", LogInToken) // This is public const variable
}
