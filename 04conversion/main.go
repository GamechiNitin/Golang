package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcom to our Pizza app")

	fmt.Println("Please, rate our Pizza from 1 to 5")

	// Taking user input

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks of rating,", input)
	fmt.Printf("Type of rating is, %T", input)

	// Add +1 rating in userinput

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Add 1 to your rating,", numRating+1)
	}

	 // Here for the conversion of String into float used [strconv] package.
	 // Here for removing space from user input used [strings] package.
	 // Also comma, err syntax with if and else is used. 
}
