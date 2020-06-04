package main

//Points to ponder
//	- Type conversion and underlying types (int is the underlying type for hotdog here)
//	- How to print type of go variable.
import "fmt"

type hotdog int

var x hotdog

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
