package main

// Points to ponder
//-  switch statments (default switch) and switch with expression/statements

import "fmt"

func main() {

	//default switch , by default
	switch {
	case true:
		fmt.Println("This is true case.")
	case false:
		fmt.Println("This is false case.")
	default:
		fmt.Println("This default case.")
	}

	favSport := "cricket"
	switch favSport {
	case "tennis":
		fmt.Println("This is Tennis Case.")
	case "cricket":
		fmt.Println("This is Cricket Case.")
		fallthrough
	case "badminton":
		fmt.Println("This is badminton case")
	default:
		fmt.Println("This Default case.")
	}
}
