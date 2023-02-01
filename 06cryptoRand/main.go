package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcom to Crypto Dice App")

	// ::::::::::::::::::::: Crypto/Rand ::::::::::::::::

	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(6))

	fmt.Println("Dice role : ", randomNumber)

}
