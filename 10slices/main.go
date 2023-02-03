package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	// Initialize Slices
	var fruitList = []string{"Apple", "Mango", "Orange"}

	fmt.Println("FruitList ", fruitList)
	fmt.Printf("FruitList type %T \n", fruitList)

	// In slice use append method to add the data to slice

	fruitList = append(fruitList, "Grapes", "Peach")
	fmt.Println("New Fruit List ", fruitList)

	// Slicing the list
	fruitList = append(fruitList[1:])
	fmt.Println("Slice Fruit List ", fruitList)

	// Slicing the list between two points
	fruitList = append(fruitList[1:3])
	fmt.Println("Slice Fruit List ", fruitList)

	//::::::::::::::::::::::::::: Using make()

	highScores := make([]int, 4) // Default Memory Allocation

	highScores[0] = 244
	highScores[1] = 544
	highScores[2] = 444
	highScores[3] = 740

	highScores = append(highScores, 999, 654, 777) // Reallocation of Memory [Append keyword]
	fmt.Println(highScores)

	//  Sort

	sort.Ints(highScores) // Sort the list in increasing order
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // Get bool Check isSorted or not
}
