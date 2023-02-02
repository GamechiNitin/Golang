package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	// Note : Length is very important, and it is always required

	var fruitList [4]string // Syntax
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Orange"
	fruitList[3] = "Grapes"
	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list length is : ", len(fruitList))

	// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

	var vegList = [3]string{"Potato", "Palak", "Peas"} // Syntax
	fmt.Println("Veg list is : ", vegList)
	fmt.Println("Veg list length is : ", len(vegList))

}
