package main

import (
	"fmt"
)

func main() {

	resultch := make(chan string)

	resultch <- "golang company" // adding data to channel
	result := <-resultch         // reading data from channel to variable

	fmt.Println(result)
}
