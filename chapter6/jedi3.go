package main

//Points to ponder
//  - Use defer keyword to show that deffered function runs after  the func containing it exits

import "fmt"

func right() {
	fmt.Println("Function right is called.")
}
func another() {
	fmt.Println("Function another is called.")
}
func main() {
	fmt.Println("Function main is called.")
	defer right()
	another()
}
