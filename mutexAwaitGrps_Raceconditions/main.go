package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Test for race_conditions")

	wg := &sync.WaitGroup{} // storing a wait_group pointer

	var data = []int{0}

	// IIFE in golang
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		data = append(data, 1)
		fmt.Println("One!")
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		data = append(data, 2)
		fmt.Println("Two!")
	}(wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		data = append(data, 3)
		fmt.Println("Three!")
	}(wg)

	wg.Wait()
}
