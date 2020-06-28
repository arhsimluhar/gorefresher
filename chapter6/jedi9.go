package main

import "fmt"

//Points to ponder
//Closure is when we have “enclosed” the scope of a variable in some code block.
//For this hands-on exercise, create a func which “encloses” the scope of a variable:
func foo() {
	{
		// This variable is only available within this parenthesis scope
		var x int = 3
		fmt.Println("Enclosed Value X:", x)
	}

}

func main() {
	foo()
}
