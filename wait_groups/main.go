package main

import (
	"fmt"
	"net/http"
)

func main() {
	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
	}

	for _, web := range websiteList {
		getStatusCode(web)
	}
}

func getStatusCode(endpoint string) {
	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops!")
	}

	fmt.Printf("%d status code endpoint for %s \n", result.StatusCode, endpoint)
}
