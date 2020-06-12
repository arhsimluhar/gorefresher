package main

//Points to ponder
//  - Anonymous Structs.
//  - How to define, use and access them.

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Rahul Mishra",
		friends: map[string]int{
			"vikram": 2,
			"suraj":  3,
		},
		favDrinks: []string{"smoothie", "cold-drinks"},
	}

	fmt.Printf("\n\nDetails of object s:\n")
	fmt.Printf("First Name: %s\n", s.first)
	for k, v := range s.friends {
		fmt.Printf("\tName: %s , Contact ID: %d\n", k, v)
	}
	fmt.Println("Favourite Drinks: ")
	for i, v := range s.favDrinks {
		fmt.Printf("\t%d %s\n", i, v)
	}

}
