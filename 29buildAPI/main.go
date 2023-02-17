package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Building Api in Golang")

	// Controller Routing - Mux
	r := mux.NewRouter()

	// Data Seeding
	courses = append(courses, Course{
		CourseId:    "3",
		CourseName:  "Flutter",
		CoursePrice: 3999,
		Author: &Author{
			FullName: "Google.com",
			Website:  "Flutter.dev",
		},
	})
	courses = append(courses, Course{
		CourseId:    "6",
		CourseName:  "Golang",
		CoursePrice: 999,
		Author: &Author{
			FullName: "Golang team",
			Website:  "go.dev",
		},
	})

	// Routing or Navigation
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// Fake DataBase
var courses []Course

// Middleware, Helper - File
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// Controller - File
func serveHome(w http.ResponseWriter, r *http.Request) {
	// Serve Home Route
	w.Write([]byte("<h1>Welcome to Build APIs in Golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course request")

	// Setting Header
	w.Header().Set("Content-Type", "application/json")

	// Encoding json
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course request")

	// Setting Header
	w.Header().Set("Content-Type", "application/json")

	// Grab the Id from Request
	params := mux.Vars(r)

	// Loop ID through courses, Find matching ID and Return response.
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	// Create one Course
	fmt.Println("Create one course")

	// Setting Header
	w.Header().Set("Content-Type", "application/json")

	// What if: body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please add some data")
	}

	// What about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Provide a valid course name")
		return
	}

	// Check Duplicate Course Title

	for _, values := range courses {
		if values.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course with same name already exit!")
			return
		}
	}

	// Generate Unique Id
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // Int to String

	// Append course into courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	// Update one Course
	fmt.Println("Update one course")

	// Setting Header
	w.Header().Set("Content-Type", "application/json")

	//Grab the request ID
	params := mux.Vars(r)

	// Loop ID through couses, Find macthing ID and Remove the Course and append new course with same ID.
	for index, value := range courses {
		if value.CourseId == params["id"] {

			// Deleting previous course data by mathcing ID
			courses = append(courses[:index], courses[index+1:]...)

			// Request Json Decoding
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			// Updating...
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: ID not found Senario
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	// Delete one course
	fmt.Println("Delete one course")

	// Setting Header
	w.Header().Set("Content-Type", "application/json")

	//Grab the request ID
	params := mux.Vars(r)

	// Loop ID through couses, Find macthing ID and Remove the Course and append new course with same ID.
	for index, value := range courses {
		if value.CourseId == params["id"] {

			// Deleting previous course data by mathcing ID
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			break
		}
	}
	//TODO: ID not found Senario
}
