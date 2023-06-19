package main

import (
	"fmt"
)

func main() {

	resultch := make(chan string, 10) // buffered channel

	resultch <- "golang company" // adding data to channel
	result := <-resultch         // reading data from channel to variable

	fmt.Println(result)
}
