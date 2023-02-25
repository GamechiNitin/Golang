package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")

	// sChannel := make(chan int) // Step 1 : Creating Simpel channel variable
	myChannel := make(chan int, 2) // Step 1 : Creating Buffer channel variable
	wg := &sync.WaitGroup{}

	// myChannel <- 5 // Step 2 : Pushing value in channel
	// fmt.Println(<-myChannel)

	wg.Add(2)
	// 1 GoRoutines -
	go func(ch <-chan int, wg *sync.WaitGroup) { // ch <-chan int - Receive OnlyChannel

		// To check weather the channed is open or not
		channelValue, isChannelOpen := <-myChannel

		fmt.Println(isChannelOpen) // Step 5 : Channels is open/closed
		fmt.Println(channelValue)  // Step 6 : Print channel value
		wg.Done()
	}(myChannel, wg)

	// 2 GoRoutines
	go func(ch chan<- int, wg *sync.WaitGroup) { // ch chan<- int - Send OnlyChannel
		myChannel <- 5
		myChannel <- 25  // Step 3 : To push value in Channel
		close(myChannel) //Step 4 : To Close the Channel
		wg.Done()
	}(myChannel, wg)

	wg.Wait()

}
