package main

import (
	"fmt"
	"net/url"
)

const urls string = "https://nitingamechi:3000/learn?coursename=golang&paymentId=dnjnddnfvkn"

func main() {
	fmt.Println("Handling URLs in Golang")
	fmt.Println(urls)

	// Parsing : Conversion from one format to another
	// Compiler Parsing : Breaking code statement into smaller pieces and analyze it.

	result, err := url.Parse(urls)
	chechNilError(err)

	fmt.Println("Scheme : ", result.Scheme)
	fmt.Println("Host : ", result.Host)
	fmt.Println("HostName : ", result.Hostname())
	fmt.Println("Path : ", result.Path)
	fmt.Println("Port : ", result.Port())
	fmt.Println("RawQuery : ", result.RawQuery) // Params

	// Better way to write RawQuery

	qParams := result.Query()
	// Type is url.Values => key,value pairs
	fmt.Println("\n::::::::::: qParams :::::::::::::")
	fmt.Printf("The type of query params is : %T\n", qParams)
	fmt.Println(qParams)

	fmt.Println("Featching values by keys :", qParams["coursename"])

	// Using Loop extracting qparams
	for key, val := range qParams {
		fmt.Printf("Key : %v and params : %v\n", key, val)
	}

	createUrl()
}

func createUrl() {

	// Never forget & sign
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "nitingamehchi.dev",
		Path:     "/developer",
		RawPath:  "user=nitin",
		RawQuery: "test=321",
	}
	generatedUrls := partsOfUrl.String()
	fmt.Println("\n::::::::::: Create a New URLs :::::::::::::")
	fmt.Println(generatedUrls)
}

// HELPER : Error handle function
func chechNilError(err error) {
	if err != nil {
		panic(err)
	}
}
