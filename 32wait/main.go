package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Package sync : WaitGroup method
var wg sync.WaitGroup // Step 1 : Declare variable

func main() {
	fmt.Println("Go routines")

	// Step 1 : Declare variable
	// Step 2 : Add go GoRoutines
	// Step 3 : Add wg.Add(1)
	// Step 4 : Add wg.Wait()
	// Step 5 : Add wg.Done()

	websiteList := []string{
		"https://loc.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web) // Step 2 : Add go GoRoutines
		wg.Add(1)             // Step 3 : For one call
	}

	wg.Wait() // Step 4 : End of main method
}

func getStatusCode(endpoint string) {
	defer wg.Done() // Step 5 : Send Exit Signal to GoRoutines
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
