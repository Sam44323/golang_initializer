/*
Marshal and Unmarshal convert a string into JSON and vice versa. Encoding and decoding convert a stream into JSON and vice versa.
*/

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Models for course and author

type Course struct {
	CourseId    string  `json="courseid"`
	CourseName  string  `json="coursename"`
	CoursePrice int     `json="price"`
	Author      *Author `json="author"`
}

type Author struct {
	Fullname string `json="fullname"`
	Website  string `json="fullname"`
}

// fake DB
var courses []Course

// middlewares

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// controllers

// serve the home_route

func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to the Home Page</h1>"))
}

func getCourses(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting all the data")
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(courses)
}

func getOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Getting one course")
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // getting the id from the url
	for _, c := range courses {
		if c.CourseId == params["id"] {
			json.NewEncoder(res).Encode(c)
			return
		}
	}
	json.NewEncoder(res).Encode("No such course was found!")
}

func createOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Creating one course")
	res.Header().Set("Content-Type", "application/json")

	// if the request data is empty
	if req.Body == nil {
		json.NewEncoder(res).Encode("No data was sent!")
		return
	}

	// when the data is {}
	var course Course
	_ = json.NewDecoder(req.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(res).Encode("No data was sent!")
		return
	}

	// generating a new id and converting it to string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // converting the integer to a string
	courses = append(courses, course)
	json.NewEncoder(res).Encode(course)
	return

	// appending the new course to the courses_slice
}

func main() {

}
