package main

import "fmt"

// points to ponder:
//	- shorthand expressions
//	- printing variables with newline characters
func main() {

	x := 42
	y := "James, Bond"
	z := true

	fmt.Printf("%d %s %t\n", x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
