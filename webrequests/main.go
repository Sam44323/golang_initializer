package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Initializing the web_requester!")
	PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	url := "http://localhost:3000/"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status_code for the response: ", response.StatusCode)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("Content of the response: ", string(content))

	/*
		Alternative method for reading a response
		var responseString string.Builder
		content, _ := ioutil.ReadAll(response.Body)
		byteCount, _ := responseString.Write(content)

		fmt.Println("Content of the response: ", responseString.String()) // printing the string content of response
	*/
}

func PerformPostJsonRequest() {
	url := "http://localhost:3000/post"

	// json fake data using ``

	requestBody := strings.NewReader(`
	{"data": "This is a fake json data"}
	`)

	response, err := http.Post(url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content of the response: ", string(content))
}
