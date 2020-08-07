package main

import (
	"fmt"
)

func main() {
	x := 242
	fmt.Printf("\n%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n\n", x)
	//or just for kicks on one line :)
	fmt.Printf("%d\t\t%b\t\t%#x", x, x, x)
}
