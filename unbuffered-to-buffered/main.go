package main

import (
	"fmt"
)

func main() {

	resultch := make(chan string) // buffered channel

	go func() {
		resultch <- "golang company" // adding data to channel using go routine
	}()
	result := <-resultch // reading data from channel to variable

	fmt.Println(result)
}
