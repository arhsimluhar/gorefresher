package main

//Points to ponder
//-  Different data types like int, string and bool
//-  returning a string after formatting a string from variables (use of fmt.Sprintf)

import "fmt"

var x = 42
var y = "james bond"
var z = true

func main() {

	s := fmt.Sprintf("%d %s %t", x, y, z)
	fmt.Println(s)
}
