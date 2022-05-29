package main

import (
	"fmt"
	"net/http"
)

func main() {

}

func getStatusCode(endpoint string) {
	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops!")
	}

	fmt.Printf("%d status code endpoint for %s \n", result.StatusCode, endpoint)
}
