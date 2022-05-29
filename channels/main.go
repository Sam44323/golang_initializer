package main

import "fmt"

func main() {
	fmt.Println("Channels in golang!")

	// data can't be sent to channel unless there is a go_routine

	channel := make(chan int) // creating a channel that handles integer

	channel <- 10 // sending data to the channel

	fmt.Println(<-channel) // receiving data from the channel
}
