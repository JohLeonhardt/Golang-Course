# Golang-Course

QUESTION: 
Hands-on exercise #1 - Video .34
1.	Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
a.	42
b.	“James Bond”
c.	true
2.	Now print the values stored in those variables using 
a.	a single print statement
b.	multiple print statements

ANSWER:------
package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
-----------

QUESTION:
Hands-on exercise #2 - Video .35
1.	Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
a.	identifier “x” type int
b.	identifier “y” type string
c.	identifier “z” type bool
2.	in func main
a.	print out the values for each identifier
b.	The compiler assigned values to the variables. What are these values called?

Answer:--------
package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
//Q: The compiler assigned values to the
//variables. What are these values called?
//A: They are called zero vales.

-------------

