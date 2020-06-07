package main

//Points to ponder
//  -  create a slice of 10 elements
//  -  print the slice and type of slice

import "fmt"

func main() {

	sliced := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range sliced {
		fmt.Printf("i: %d, v: %d\n", i, v)
	}
	fmt.Printf("%T\n", sliced)

}
