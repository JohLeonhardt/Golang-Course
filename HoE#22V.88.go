package main

import (
	"fmt"
)

func main() {
	x := [5]int{10, 11, 12, 13, 45}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
