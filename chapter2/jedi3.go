package main

//Points to ponder
//   - const variables ( typed and untyped )

import "fmt"

const (
	a     = 42
	b int = 42
)

func main() {
	fmt.Println(a, b)
}
