package main

//Points to ponder
//-  var keyword and how to use it.
//-  printing variables in a single print statement.
//

import "fmt"

var x int
var y string
var z bool

func main() {

	fmt.Print(x, y, z)

	fmt.Print(x)
	fmt.Print(y)
	fmt.Print(z)

}
