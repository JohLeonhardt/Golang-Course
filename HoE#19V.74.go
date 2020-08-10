package main

import (
	"fmt"
)

func main() {
	switch {
	case true:
		fmt.Println("This should print")
	case false:
		fmt.Println("This should not print")
	}

}
