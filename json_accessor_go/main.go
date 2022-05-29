package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Platform string
	Price    int
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Initializing the JSON accessor!")
	EncodeJson()
}

// encoding any data into JSON
func EncodeJson() {

	courses := []course{
		{
			"Go",
			"Golang",
			1000,
			"123",
			[]string{"Go", "Golang", "Programming"},
		},
		{
			"Py",
			"Python",
			2000,
			"456",
			[]string{"Python", "Programming"},
		},
		{
			"JS",
			"Javascript",
			3000,
			"789",
			nil,
		},
	}

	// packaging this data as JSON

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}
