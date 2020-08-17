package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)

	ii2 := []int{4, 5, 6, 7, 8}
	n2 := bar(ii2)
	fmt.Println(n2)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func bar(xi2 []int) int {
	total2 := 0
	for _, v2 := range xi2 {
		total2 += v2
	}
	return total2
}
