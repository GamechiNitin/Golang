package main

import "fmt"

func main() {
	fmt.Println("Welcome to delete slices from list")

	var coursesList = []string{"Flutter", "Dart", "AWS", "Golang"}
	fmt.Println("Courses List : ", coursesList)

	var index int = 2
	coursesList = append(coursesList[:index], coursesList[index+1:]...)
	fmt.Println("Deleted Courses List : ", coursesList)

}
