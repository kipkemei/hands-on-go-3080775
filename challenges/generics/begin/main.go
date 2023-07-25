// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring
func printAny[T any](v T) { fmt.Println(v) }

// Part 2 sum function refactoring
type numeric interface {
	constraints.Float | constraints.Integer
}

// sum sums a slice of any type
func sum[T numeric](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	// call generic printAny function for each value above
	printAny("Hello")
	printAny(42)
	printAny(true)

	// call generics sum function
	fmt.Println(sum(1, 2, 3))
}
