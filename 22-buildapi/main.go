package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Models -file
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware or helper -file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

}

// controllers -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the First Ever Web Server</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request body must not be empty")
		return
	}

	// if the body is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	// generate a unique id, string
	// append into the courses
	num := rand.Int()
	course.CourseId = strconv.Itoa(num)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
