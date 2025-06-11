// Witness the least applicable use of a pointer ✨
// Pointers are refernces in Golang Anyhow

package main

import "fmt"

func main() {
	var i int64 = 134
	p := &i
	fmt.Printf("Reading i through pointer p: %v \n", *p)
	for count := 0; count < 10; count++ {
		*p += 1
		fmt.Printf("Adding 1 though pointer: %v \n", *p)
	}
}
