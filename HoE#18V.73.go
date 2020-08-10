package main

import (
	"fmt"
)

func main() {
	x := 42
	if x == 40 {
		fmt.Println("The value was 40")
	} else if x == 41 {
		fmt.Println("The value not 40")
	} else {
		fmt.Println("Then the value was 41 or 42")
	}
}
