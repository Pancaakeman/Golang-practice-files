package main

import "fmt"

func main(){
	silly := 0

	for i := 0; i < 10; i++ {
		silly += i
		fmt.Printf("i = %v ", i)
		fmt.Printf("Silly = %v \n", silly)
	}
}