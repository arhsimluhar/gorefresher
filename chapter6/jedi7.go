package main

//Points to ponder
// - Create a func which returns a func
// - Assign the returned func to a variable
// - Call the returned func

import "fmt"

func foo() func() int {
	return func() int {
		return 42
	}
}

func main() {
	bar := foo()
	fmt.Println("Calling the returned function", bar())
}
