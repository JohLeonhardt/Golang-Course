package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("First foo")
}

func bar() {
	fmt.Println("Second bar")
}
