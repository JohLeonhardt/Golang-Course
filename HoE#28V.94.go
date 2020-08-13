package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(mp)
	x := [][]string{jb, mp}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println("Records:", i)
		for j, k := range v {
			fmt.Printf("\t\tIndex Postion: %v\tValue: %v\n", j, k)
		}

	}
}
