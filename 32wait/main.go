package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Package sync : WaitGroup method
var wg sync.WaitGroup // Step 1 : Declare wait pointer variable
var mutex sync.Mutex  // Step 1 : Declare Mutex pointer variable

var signals = []string{"test"}

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
	fmt.Println(signals)

}

func getStatusCode(endpoint string) {
	defer wg.Done() // Step 5 : Send Exit Signal to GoRoutines
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		// ::::::::::: Mutex :::::::::::::::
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		// ::::::::::: Mutex :::::::::::::::
		
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
