package main

import "fmt"

var pow = []int{0, 1, 4, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}
}
