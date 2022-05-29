package main

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

func main() {

}
