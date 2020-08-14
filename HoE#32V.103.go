package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "James",
		last:       "Bond",
		favFlavors: []string{"Vanilla", "Cherry", "Bubble-gum"},
	}
	p2 := person{
		first:      "Naomi",
		last:       "McArthur",
		favFlavors: []string{"Banana", "Wallnut", "Apple"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println("\t", i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for g, h := range p2.favFlavors {
		fmt.Println("\t", g, h)
	}
}
