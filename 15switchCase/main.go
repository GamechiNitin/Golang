package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

    // Generating Random Number 
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is : ", randomNumber)
	
	switch randomNumber {
	case 1 :
		fmt.Println("Dice value is 1 you can open")
	case 2 :
		fmt.Println("You can move 2 spot")
	case 3 :
		fmt.Println("You can move 3 spot")
		fallthrough // Fall through Next Case
	case 4 :
		fmt.Println("You can move 4 spot")
	case 5 :
		fmt.Println("You can move 5 spot")
	case 6 :
		fmt.Println("You can move 6 spot and roll dice again")
	default :
		fmt.Println("What was that!")
	}
}
