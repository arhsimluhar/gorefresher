package main

//Points to ponder
//  - Create and use anonymous functions

import "fmt"

func main() {
	square := func(x int) int {
		return x * x
	}
	fmt.Println("Square of 4 is", square(4))
}
