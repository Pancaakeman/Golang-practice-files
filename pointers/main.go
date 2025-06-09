package main

import "fmt"

func main() {
	var i int64 = 134
	p := &i
	fmt.Printf("Reading i through pointer p: %v \n", *p)
	for count := 0; count < 10; count++ {
		*p += 1
		fmt.Printf("Added 1 to i using pointer: %v \n", *p)
	}
}
