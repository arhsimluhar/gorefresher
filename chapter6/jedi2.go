package main

//Points to ponder
//  - Create a func with identifier foo that
// 		- takes in a variadic parameter of type int
// 		- pass in a value of type []int into your func (unfurl the []int)
// 		- returns the sum of all values of type int passed in
//  - Create a func with identifier bar that
//    - takes in a parameter of type []int
//    - returns the sum of all values of type int passed in

import "fmt"

func foo(ii ...int) int {

	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

func bar(ii []int) int {

	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	sum := foo(x...)
	fmt.Printf("Sum of Slice %d.\n", sum)
	sum = bar(x)
	fmt.Printf("Sum of Slice %d.\n", sum)

}
