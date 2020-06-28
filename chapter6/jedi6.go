package main

//Points to ponder
//  - Create and use anonymous functions
//  - Assign a function to a variable and call it

import (
	"fmt"
	"math"
)

func main() {
	func(x int) {
		fmt.Printf("Square of %d is %d.\n", x, x*x)

	}(4)
	cube := func(x float64) float64 {
		return math.Pow(x, 3)
	}
	fmt.Println("Cube of 4 is :", cube(4))
}
