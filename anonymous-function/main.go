package main

import (
	"fmt"
	"time"
)

func main() {

	go func() { ///anonymous function

		fetchResource() //async call
		fetchResource()
		fetchResource()
		fetchResource()
		fetchResource()
		fetchResource()
		fetchResource()
	}()
	fmt.Println("I am complete")
}

func fetchResource() string {

	time.Sleep(time.Second * 2)

	return "some string"
}
