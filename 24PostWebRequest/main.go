package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post Web Request in Golang")
	PerformPostRequest()
}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"

	// Fake payload
	payload := strings.NewReader(`
		{
			"coursename":"Let's go with Golang",
			"price":0,
			"platfrom":"nitingamechi.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", payload)

	chechNilError(err)
	defer response.Body.Close()
	fmt.Println("Post response : ", response.Status)

	var responseBuilder strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	chechNilError(err)

	byteCount, _ := responseBuilder.Write(content)
	fmt.Println("ByteCount", byteCount)
	fmt.Println("Response Builder O/P : ", responseBuilder.String())
}

func chechNilError(err error) {
	if err != nil {
		panic(err)
	}
}
