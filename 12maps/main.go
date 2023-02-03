package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languagesList := make(map[string]string)
	languagesList["Js"] = "Javascript"
	languagesList["go"] = "Golang"
	languagesList["py"] = "Python"
	languagesList["RB"] = "Ruby"

	fmt.Println("List of maps : ", languagesList)

	// Get Individual key value
	fmt.Println("Key : ", languagesList["go"])

	// Delete Individual key value
	delete(languagesList, "py")
	fmt.Println("Deleted List of maps : ", languagesList)

	// Loops with maps
	for key, value := range languagesList {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
