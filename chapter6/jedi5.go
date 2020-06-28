package main

//Points to ponder
// - create a type SQUARE
// - create a type CIRCLE
// - attach a method to each that calculates AREA and returns it
// 		circle area= π r 2
// 		square area = L * W
// - create a type SHAPE that defines an interface as anything that has the AREA method
// - create a func INFO which takes type shape and then prints the area
// - create a value of type square
// - create a value of type circle
// - use func info to print the area of square
// - use func info to print the area of circle

import (
	"fmt"
	"math"
)

type SQUARE struct {
	side float32
}
type CIRCLE struct {
	radius float32
}

type SHAPE interface {
	Area() float32
}

func (q SQUARE) Area() float32 {
	return q.side * q.side
}

func (c CIRCLE) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func INFO(x SHAPE) {
	fmt.Printf("Area of %T is %f.\n", x, x.Area())
}

func main() {
	square := SQUARE{side: 4}
	circle := CIRCLE{radius: 4}
	INFO(square)
	INFO(circle)
}
