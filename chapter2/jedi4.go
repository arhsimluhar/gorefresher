package main

import "fmt"

func main() {
	var x int = 4
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

	x = x << 1

	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

}
