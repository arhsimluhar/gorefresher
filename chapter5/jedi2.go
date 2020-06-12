package main

//Points to ponder
//  - Named structs
//  - Using maps to store structs and accessing them though for loop.

import (
	"fmt"
)

type person struct {
	firstName       string
	lastName        string
	iceCreamflavors []string
}

func main() {
	rahul := person{
		firstName:       "Rahul",
		lastName:        "Mishra",
		iceCreamflavors: []string{"mango", "butterscotch", "chocolate"},
	}
	shyam := person{
		firstName:       "Shyam",
		lastName:        "Bro",
		iceCreamflavors: []string{"vanilla", "butterscotch"},
	}

	data := map[string]person{"Mishra": rahul, "Bro": shyam}

	for k, v := range data {
		fmt.Printf("Record entry for %s\n", k)
		fmt.Printf("First Name: %s\n", v.firstName)
		fmt.Printf("Last Name: %s\n", v.lastName)

		fmt.Println("Favorite Flavors of Ice-Cream:")
		for index, flavour := range v.iceCreamflavors {
			fmt.Printf("%d %s\n", index, flavour)
		}
		fmt.Println()
	}

}
