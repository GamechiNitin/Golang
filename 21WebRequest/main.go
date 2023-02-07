package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://loc.dev"

func main() {
	fmt.Println("Web Request in Golang")

	response, err := http.Get(url)

	chechNilError(err)

	fmt.Printf("Response type is : %T\n", response)
	fmt.Println("Status Code is : ", response.StatusCode)
	fmt.Println("Status is : ", response.Status)
	fmt.Println("Status is : ", response.Request)

	defer response.Body.Close() // Developer Responsibility to close it.
	readResponse(response)
}

func readResponse(response *http.Response) {
	// Majority of reading is done by io package
	dataByte, err := io.ReadAll(response.Body) // ReadResponse
	chechNilError(err)
	content := string(dataByte)
	fmt.Println("Databyte is : ", string(content))
}

// HELPER : Error handle function
func chechNilError(err error) {
	if err != nil {
		panic(err)
	}
}
