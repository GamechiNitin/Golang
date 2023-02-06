package main

import "fmt"

func main() {

	// Defer keyword add the nextline code in Stack and execute in LIFO afterwards.
	defer fmt.Print(" World")
	defer fmt.Print(" one")
	defer fmt.Print(" from")
	
	fmt.Print("Hello ")
	myDefer()
}


// First Stact :-  World, one, from o/p -> from one world
// Second Stact :- 01234 o/p -> 43210
// Output :- Hello 43210 from one world

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
