package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	weekDays := []string{"Sunday", "Monday", "Tuesday", "Thursday", "Friday", "Saturday"}

	// Print the List
	fmt.Println(weekDays)

	// For Loop
	for i := 0; i < len(weekDays); i++ {
		fmt.Println("Loop :", weekDays[i]) // i => Return the index
	}

	// For Range Loop
	for r := range weekDays {
		fmt.Println("Range : ", weekDays[r]) // i => Return the index
	}

	// For Range Loop return Index and Value
	for index, value := range weekDays {
		fmt.Printf("Index is %v and value is %v\n", index, value)
	}

	roughValue := 1
	for roughValue < 10 { // Like a While loop

		// Continue : Skip over one phase. And we have to provide that phase.
		if roughValue == 3 {
			roughValue++
			continue
		}

		// Goto : goto and label statements which are used to jump from one point to other
		if roughValue == 6 {
			goto nitinGamechi
		}
		// Break : Terminate the loop all together.
		if roughValue == 7 {
			break
		}

		fmt.Println("Value : ", roughValue)
		roughValue++
	}

nitinGamechi:
	fmt.Println("Jumping at nitingamechi.in")
}
