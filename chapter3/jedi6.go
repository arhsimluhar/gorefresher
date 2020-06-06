package main

// Points to ponder
//  - If-else-if ladder

import "fmt"

func main() {
	x := 11
	if x%2 == 0 {
		fmt.Println("The number is even.")
	} else if x%2 == 1 {
		fmt.Println("The number is odd.")
	} else {
		fmt.Println("Should never reach this condition.")
	}

}
