package main

//Points to ponders
//  - Embedded structs.
//  - How to define them and access them.

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourwheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck_kun := truck{
		vehicle: vehicle{
			color: "Red",
			doors: 2,
		},
		fourwheel: true,
	}

	sedan_kun := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println("Specifications for Truck kun:")
	fmt.Println(truck_kun.color)
	fmt.Println(truck_kun.doors)
	fmt.Println(truck_kun.fourwheel)

	fmt.Println("Specifications for Sedan kun:")
	fmt.Println(sedan_kun.color)
	fmt.Println(sedan_kun.doors)
	fmt.Println(sedan_kun.luxury)

	fmt.Printf("%#v\n", truck_kun)
	fmt.Printf("%#v\n", sedan_kun)

}
