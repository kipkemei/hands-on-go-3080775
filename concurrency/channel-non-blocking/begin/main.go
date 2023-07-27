// concurrency/channel-non-blocking/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// declare two time channels to read from
	timeChan1 := time.After(200 * time.Millisecond)
	timeChan2 := time.After(400 * time.Millisecond)

	// use a for loop to iterate indefinitely until channels return a value
	for {
		select {
		case <- timeChan1:
			fmt.Println("timeChan1")
			return
		case <- timeChan2:
			fmt.Println("timeChan2")
			return
		default:
			fmt.Println("default")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
