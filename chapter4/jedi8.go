package main

//Points to ponder
//  - Using maps
//  - Using nested for loops with range to print map elements.

import "fmt"

func main() {

	person := map[string][]string{}
	person["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	person["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	person["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}

	fmt.Println(person)

	for key, _ := range person {
		fmt.Printf("This is record for %s:\n", key)
		for i, v := range person[key] {

			fmt.Printf("Index: %d , item: %s\n", i, v)
		}
		fmt.Println("")
	}

}
