package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition in Golang")

	var score = []int{0}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	wg.Add(3) // Add Goroutines one by one or All Together
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mutex.Lock()
		score = append(score, 1)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mutex.Lock()
		score = append(score, 2)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mutex.Lock()
		score = append(score, 3)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(score)
}
