package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Test for race_conditions")

	wg := &sync.WaitGroup{} // storing a wait_group pointer

	var data = []int{0}

	wg.Add(3) // total_number of go-routines
	// IIFE in golang
	go func(wg *sync.WaitGroup) {
		data = append(data, 1)
		fmt.Println("One!")
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		data = append(data, 2)
		fmt.Println("Two!")
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		data = append(data, 3)
		fmt.Println("Three!")
		wg.Done()
	}(wg)

	wg.Wait()
}
