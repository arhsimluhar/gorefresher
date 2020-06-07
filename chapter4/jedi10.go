package main

//Points to ponder
//  - deleting an element in map

import (
	"fmt"
)

func main() {

	person := map[string][]string{}
	person["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	person["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	person["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}

	person["rahul"] = []string{"mysterious", "talented", "dependable"}

	delete(person, "bond_james")

	for key, _ := range person {
		fmt.Printf("name: %s , attributes: %v\n", key, person[key])
	}
}
