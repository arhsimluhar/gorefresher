package main

//Points to ponder
//   - For loop with any(init, condition or post statement)

import "fmt"

func main() {
	start := 1993
	for {
		if start > 2020 {
			break

		}
		fmt.Println(start)
		start++
	}
}
