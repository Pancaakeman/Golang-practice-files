// Behold, My crappy implementation of a Square-root function
// Struggles with bigger numbers (I have no idea how to check the 2nd last iterated value)
// Future me: Maybe Make this actually work for bigger numbers
// That's about it
package main

import (
	"fmt"
)

func Sqrt(num float64) float64 {
	iterative := 1.0
	for count := 0; count < 10; count++ { // Initalises Count, Iterates until it reaches the 10 value
		iterative = iterative - (iterative*iterative-num)/(2*iterative) //silly formula to guess a sqrt
		fmt.Printf("Guess: %v \n", iterative)
	}
	return iterative //returns the last value reached
}

func main() {
	defer fmt.Println(Sqrt(9))
	defer fmt.Println("Square Root Should be: ")
}
