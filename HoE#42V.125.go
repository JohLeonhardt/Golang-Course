package main

import (
	"fmt"
)

func main() {
	fmt.Println("test1")

	f := func() {
		fmt.Println("func print")
	}
	f()
}
