package main

import (
	"fmt"
)

const (
	y1 = iota + 2021
	y2
	y3
	y4
)

func main() {
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
	// or in one line
	fmt.Println(y1, y2, y3, y4)
}
