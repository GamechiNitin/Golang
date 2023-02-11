package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("PostForm Web Request in Golang")
	PerfromPostFormRequest()
}

func PerfromPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "Nitin")
	data.Add("lastName", "Gamechi")
	data.Add("email", "nitingamechi@gmail.com")

	response, err := http.PostForm(myUrl, data)
	checkNilError(err)
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
