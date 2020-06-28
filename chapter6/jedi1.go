package main

//Points to ponder
//  - create a func with identifier foo that returns int
//  - create a func with identifier bar that returns int and string
//  - call both the func
//  - print out their results

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 33, "Hello world!\n"
}

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a)
	fmt.Println(b, c)

}
