package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("Welcome to mongo-db in Golang")

	// Required => 1. Change Url(connectionString) 2. Run : go build main.go
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
