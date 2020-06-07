package main

//Points to ponder
//  - Two dimensional slices
//  - Using nested loops to print them

import "fmt"

func main() {
	var x [][]string
	first := []string{"James", "Bond", "Shaken, not stirred"}
	second := []string{"Miss", "Moneypenny", "Hello, James"}

	x = append(x, first, second)
	fmt.Println(x)

	for _, arr := range x {
		for _, item := range arr {

			fmt.Println(item)
		}

	}
}
