package main

//Points to ponder
//  - Using composite literal define a array
//  - print the array values using range.

import "fmt"

func main() {
	num := [5]int{0, 1, 2, 3, 4}

	for index, value := range num {
		fmt.Printf("Index: %d, Value: %d.\n", index, value)

	}
	fmt.Printf("%T", num)
}
