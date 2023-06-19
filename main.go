//Parallelism vs concurrency

package main

import (
	"fmt"
	"time"
)

func main() {

	fetchResource()
	fetchResource()
	fetchResource()
	fetchResource()
	fetchResource()
	fetchResource()
	fmt.Println("I am complete")
}

func fetchResource() string {

	time.Sleep(time.Second * 2)

	return "some string"
}

//After execute the code please commented all the above code and uncomment the below commented code and execute

////concurrency
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	go fetchResource() //async call
// 	go fetchResource()
// 	go fetchResource()
// 	go fetchResource()
// 	go fetchResource()
// 	go fetchResource()
// 	go fetchResource()
// 	fmt.Println("I am complete")
// }

// func fetchResource() string {

// 	time.Sleep(time.Second * 2)

// 	return "some string"
// }
