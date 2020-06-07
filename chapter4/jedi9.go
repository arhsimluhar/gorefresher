package main

//Points to ponder
//  -  Adding an element to map

import "fmt"

func main() {

	person := map[string][]string{}
	person["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	person["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	person["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}

	person["rahul"] = []string{"mysterious", "talented", "dependable"}

	fmt.Println(person)
}
