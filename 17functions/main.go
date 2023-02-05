package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")

	getter()

	result := doAddition(5, 9)
	fmt.Println("Result is : ", result)

	proResult, message := proAddition(5, 9, 45, 2, 87, 44)
	fmt.Println("Pro Result is : ", proResult)
	fmt.Println("Pro Message is : ", message)
}

// Simple function
func getter() {
	fmt.Println("Namastey, Jai hind dasto")
}

// function which return value
func doAddition(v1 int, v2 int) int {
	return v1 + v2
}

// function which multiple return value and infinity(multiple) parameters
func proAddition(values ...int) (int, string) {
	// ... verdiact function

	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi this is a pro result function."
}
