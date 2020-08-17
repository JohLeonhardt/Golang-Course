package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (s person) speak() {
	fmt.Println("`I am", s.first, s.last, "and I am", s.age, "years old.`")
}

func main() {
	sa1 := person{
		first: "Bruce",
		last:  "Banner",
		age:   39,
	}
	sa1.speak()
}
