package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Dice App ")

	// There are two way to generate random number. 1) rand 2) Crypto Package

	// ::::::::::::::::::::: Math/Rand ::::::::::::::;;

	// Random Seed Number with the help of time package
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Dice roll : ", rand.Intn(6)+1)

}
