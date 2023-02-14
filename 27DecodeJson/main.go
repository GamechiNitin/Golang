package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Decoding Json in Golang")
	DecodeJson()
}

// Struct : Part 2 - Struct Alias
type courseName struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// JSON Decoding
func DecodeJson() {
	tempJson := []byte(`
	{
		"coursename": "Flutter and Dart",
		"Price": 299,
		"website": "nitingamechi.in",
		"tags": [
				"Mobile App Development",
				"Android Development"
		]
	}
	`)

	// Checking Json is Valid of Not
	checkJson := json.Valid(tempJson)

	// First way : with the help of Struct
	fmt.Println("\n::::::::::::: First Way ::::::::::")
	var course courseName
	if checkJson {
		fmt.Println("Json is valid")
		json.Unmarshal(tempJson, &course)
		fmt.Printf("%#v\n", course)

	} else {
		fmt.Println("JSON IS NOT VALID")
	}

	// Second Way : Key and Value Map
	fmt.Println("\n::::::::::::: Second Way ::::::::::")
	var onlineData map[string]interface{}
	json.Unmarshal(tempJson, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	fmt.Println("\n::::::::::::: For Loop ::::::::::")
	for k, v := range onlineData {
		fmt.Printf("Key : %v, Value : %v, Type : %T\n", k, v, v)
	}

}
