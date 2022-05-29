package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/Sam44323/fs_API/routers"
)

func main() {
	fmt.Printf("MongoDB API initiating!")
	fmt.Println("Initiating the server!")
	log.Fatal(http.ListenAndServe(":3000", routes.Router()))
}
