package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
		age       int
	}{
		first: "James",
		friends: map[string]int{
			"Naomi":   118789,
			"Hulk":    999999,
			"MrWayne": 12343,
		},
		favDrinks: []string{
			"Martini",
			"Hulkjuice",
			"Batjuice",
		},
		age: 31,
	}
	fmt.Println(p1.first)
	fmt.Println(p1.friends)
	fmt.Println(p1.favDrinks)
	fmt.Println(p1.age)

	for k, v := range p1.friends {
		fmt.Println(k, v)
	}
	for i, val := range p1.favDrinks {
		fmt.Println(i, val)
	}

}
