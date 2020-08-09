package main

import (
	"fmt"
)

func main() {
	x := 1981
	for {
		if x > 2020 {
			break
		}
		fmt.Println(x)
		x++
	}
}
