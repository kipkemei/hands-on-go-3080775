// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var i any = 1

	// perform a type assertion, handle the error and print out the asserted string value of i

	if _, ok := i.(int); !ok {
		fmt.Printf("%v is not an int\n", i)
	} else {
		fmt.Printf("%v, %T\n", i.(int), i.(int))
	}
}
