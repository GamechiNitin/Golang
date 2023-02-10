package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Get Web Request in Golang")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	chechNilError(err)

	defer response.Body.Close()

	fmt.Println("Status Code is : ", response.StatusCode)
	fmt.Println("Resposne ContentLength is : ", response.ContentLength)
	fmt.Println("Request Url is : ", response.Request.URL)

	// // First Way
	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("Normal Way:", string(content))

	// Another Way
	var responseBuilder strings.Builder
	anotherContent, err := ioutil.ReadAll(response.Body)
	chechNilError(err)
	byteCount, err := responseBuilder.Write(anotherContent)
	chechNilError(err)
	fmt.Println("Another Response : ", byteCount)
	fmt.Println("Response Builder O/P : ", responseBuilder.String())

}

// HELPER : Error handle function
func chechNilError(err error) {
	if err != nil {
		panic(err)
	}
}
