package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/Sam44323/fs_API/routers"
)

func main() {
	fmt.Printf("Server is active \n")
	log.Fatal(http.ListenAndServe(":3000", routes.Router()))
}
