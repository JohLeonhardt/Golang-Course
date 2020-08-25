package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (b ByLast) Len() int           { return len(b) }
func (b ByLast) Swap(k, g int)      { b[k], b[g] = b[g], b[k] }
func (b ByLast) Less(k, g int) bool { return b[k].Last < b[g].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Elbert",
		Last:  "Astone",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "Sarah",
		Last:  "Camalot",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	fmt.Println("***  Original output without Sort besides Sort by Saying (Alphabetical) ***")

	//we could change the "u" to "user" but I left it at "u" for readability
	users := []user{u1, u2, u3}
	for i, u := range users {
		fmt.Println("User #", i)
		fmt.Println("\t", u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, v := range u.Sayings {
			fmt.Println("\t\t\t", v)
		}
	}
	fmt.Println("****** Sorted by Age (Ascending) & Sort by Saying (Alphabetical) *********")

	sort.Sort(ByAge(users))
	for i, u := range users {
		fmt.Println("User #", i)
		fmt.Println("\t", u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t\t\t", v)
		}
	}
	fmt.Println("*** Sorted by Last Name (Alphabetical) & Sort by Saying (Alphabetical) ***")

	sort.Sort(ByLast(users))
	for i, u := range users {
		fmt.Println("User #", i)
		fmt.Println("\t", u.First, u.Last, u.Age)
		for _, v := range u.Sayings {
			fmt.Println("\t\t\t", v)
		}
	}
}
