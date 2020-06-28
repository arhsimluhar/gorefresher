package main

//Points to ponder
// A “callback” is when we pass a func into a func as an argument. For this exercise,
// 		pass a func into a func as an argument

import "fmt"

func circle() {
	fmt.Println("This is a circle.")
}

func square() {
	fmt.Println("This is a square.")
}

func foo(x func()) {
	x()
}

func main() {
	foo(circle)
	foo(square)
}
