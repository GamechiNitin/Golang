package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Json in Golang")
	EncodeJson()
}

// Struct : Part 2 - Struct Alias
type courseName struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// JSON Encoding
func EncodeJson() {
	newCourses := []courseName{
		{
			"Flutter and Dart",
			299,
			"nitingamechi.in",
			"abc321",
			[]string{"Mobile App Development", "Android Development"},
		},
		{
			"Golang",
			399,
			"nitingamechi.in",
			"kbc321",
			[]string{"Web Development", "Backend Development"},
		},
		{
			"Node js",
			199,
			"nitingamechi.in",
			"lmn321",
			nil,
		},
	}

	// Marshal Example
	fmt.Println("\n:::::::::::::::::::::::::: Json Marshal :::::::::::::::::::")
	finalJson, err := json.Marshal(newCourses)
	checkNilError(err)
	fmt.Printf("%s\n", finalJson)

	// Marshal Indent Example
	fmt.Println("\n:::::::::::::::::::::::::: Json MarshalIndent :::::::::::::::::::")
	finalIndentJson, err := json.MarshalIndent(newCourses, "", "\t")
	checkNilError(err)
	fmt.Printf("%s\n", finalIndentJson)
}

// Helper error
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
