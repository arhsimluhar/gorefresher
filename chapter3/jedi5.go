package main

//Points to ponder
//   - modulo operator (%)

import "fmt"

func main() {

	start := 10
	for {
		if start > 100 {
			break
		}
		fmt.Printf("%d %% 4  = %d \n", start, start%4)
		start++
	}
}
