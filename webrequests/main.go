package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Initializing the web_requester!")
	PerformGetRequest()
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
}
