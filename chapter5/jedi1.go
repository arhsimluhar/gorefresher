package main

//Points to ponder
//  - Named struct
//  - Initializing struct and accessing them and printing them.

import "fmt"

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

	fmt.Printf("First Name: %s\n", rahul.firstName)
	fmt.Printf("Last Name: %s\n", rahul.lastName)
	fmt.Println("Favourite Ice-Cream flavours.")
	for _, val := range rahul.iceCreamflavors {
		fmt.Println(val)
	}

	fmt.Printf("\nFirst Name: %s\n", shyam.firstName)
	fmt.Printf("Last Name: %s\n", shyam.lastName)
	fmt.Println("Favourite Ice-Cream flavours.")
	for _, val := range shyam.iceCreamflavors {
		fmt.Println(val)
	}

}
