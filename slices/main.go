package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} // An Array

	var s []int = primes[0:6]
	s[7] = 17
	fmt.Println(s)
}
