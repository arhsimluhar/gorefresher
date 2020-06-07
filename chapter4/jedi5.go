package main

import "fmt"

//Points to ponder
//  - Delete elements from slice using append and slicing

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
