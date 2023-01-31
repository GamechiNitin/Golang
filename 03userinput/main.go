package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMessage := "Welcome to UserInput Program"
	fmt.Println(welcomeMessage)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for Pizza")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating :", input)
	fmt.Printf("Type of this rating is %T", input)
}
