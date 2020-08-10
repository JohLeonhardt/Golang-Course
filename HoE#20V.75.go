package main

import (
	"fmt"
)

func main() {
	favSport := "Diving"
	switch favSport {
	case "skiing":
		fmt.Println("In pouder snow")
	case "shark diving":
		fmt.Println("in Fiji waters")
	case "surfing":
		fmt.Println("Phia beach, Auckland")
	case "basketball":
		fmt.Println("back ally")
	case "Diving":
		fmt.Println("Underwater world")
	}
}
