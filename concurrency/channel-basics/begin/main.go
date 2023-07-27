// concurrency/channels/begin/main.go
package main

import (
	"fmt"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<-int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// initialize channel for communication between main and sum goroutines
	ch := make(chan int)

	// invoke the sum function as a goroutine
	go sum(nums, ch)
	result := <- ch
	fmt.Println("Result:", result)

	// initialize a buffered channel of size 2
	ch2 := make(chan string, 2)

	// send string values into the channel
	ch2 <- "Sang"
	ch2 <- "Kipkemei"

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
