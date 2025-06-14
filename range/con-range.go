package main

import "fmt"

func test() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
