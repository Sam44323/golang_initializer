/*
Marshal and Unmarshal convert a string into JSON and vice versa. Encoding and decoding convert a stream into JSON and vice versa.
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	return c.CourseId == "" && c.CourseName == ""
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

func main() {

}
