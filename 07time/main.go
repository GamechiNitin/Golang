package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study of Golang")

	// Present time
	currentTime := time.Now()
	fmt.Println(currentTime)
	
	// Format Time : [01-02-2006 15:04:05 Monday] Docs
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))


	// Creating Time & Date
	createDate := time.Date(2025, time.September, 28, 17, 30, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
