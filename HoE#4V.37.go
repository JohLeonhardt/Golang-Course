package main

import (
	"fmt"
)

type Joh int

var x Joh

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
