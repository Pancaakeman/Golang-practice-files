package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("You are running this on:")

	switch runtime.GOOS {
	case "darwin":
		fmt.Println("MacOS")

	case "linux":
		fmt.Println("Linux")

	default:
		fmt.Printf("You are on %v", runtime.GOOS)
	}
}
