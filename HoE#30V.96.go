package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m[`pussy_galore`] = []string{`Aviation`, `Gun`, `Sexy dresses`}
	for i, v := range m {
		fmt.Println("This is the record of:", i)
		for j, k := range v {
			fmt.Println("\t", j, k)
		}
	}
}
