// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// simple greet function
func greet() string {
	return "Hola"
}

// greetWithName returns a greeting with the name
func greetWithName(name string) string {
	return "Hola " + name + "!"
}

// greetWithNameAndAge returns a greeting with the name and age of the person
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hola, my name is " + name + " and I am " +strconv.Itoa(age) + " years old"
	return
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
func divide(firstNumber, secondNumber int) (int, error) {
	if secondNumber == 0 {
		return 0, errors.New("cannot divide by 0")
	} else {
		result := firstNumber / secondNumber
		return result, nil
	}
}

func main() {
	// invoke greet function
	// fmt.Println(greet())

	// invoke greetWithName function
	// fmt.Println(greetWithName("Toni"))
	// fmt.Println(greetWithNameAndAge("Toni", 15))

	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(5, 0))
}
