package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in Golang")

	var fileName string = "./myNitinFile.txt"

	// This is File content
	content := "This is a content file - nitingamechi.in"

	file, err := os.Create(fileName) // FileName
	chechNilError(err)

	length, err := io.WriteString(file, content) // File creating
	chechNilError(err)

	fmt.Println("Length is : ", length)
	defer file.Close()

	// To Read a file
	readFile(fileName)
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName) // ReadFile
	chechNilError(err)
	fmt.Println("Databyte is : ", string(dataByte))
}

// HELPER : Error handle function
func chechNilError(err error) {
	if err != nil {
		panic(err)
	}
}
