// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
type author struct {
	firstName, lastName string
}

func main() {
	// intialize author
	a := author{
		firstName: "dennis", 
		lastName: "sang",
	}

	// print the author
	fmt.Printf("%#v\n", a)
}
