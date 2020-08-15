package main

import (
	"fmt"
)

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	Tr := truck{
		vehicle: vehicle{
			door:  4,
			color: "black",
		},
		fourWheel: true,
	}
	Se := sedan{
		vehicle: vehicle{
			door:  2,
			color: "silver",
		},
		luxury: false,
	}
	fmt.Println(Tr)
	fmt.Println(Se)

	fmt.Println(Tr.door, Tr.color, Tr.fourWheel)
	fmt.Println(Se.door, Se.color, Se.luxury)

}
