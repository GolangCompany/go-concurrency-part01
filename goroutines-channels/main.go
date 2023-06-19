package main

import (
	"fmt"
	"time"
)

func main() {

	go func() { ///anonymous function

		result := fetchResource()
		fmt.Println(result)
	}()
}
func fetchResource() string {

	time.Sleep(time.Second * 2)

	return "some string"
}
