# Golang-Course
------------------
------------------

N. Level 1 - QUESTION: 1
HoE#1V.34.go
Hands-on exercise #1 - Video .34

1.	Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
a.	42
b.	“James Bond”
c.	true
2.	Now print the values stored in those variables using 
a.	a single print statement
b.	multiple print statements

-----------

N. Level 1 - QUESTION: 2
HoE#2V.35.go
Hands-on exercise #2 - Video .35

1.	Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
a.	identifier “x” type int
b.	identifier “y” type string
c.	identifier “z” type bool
2.	in func main
a.	print out the values for each identifier
b.	The compiler assigned values to the variables. What are these values called?

-------------

N. Level 1 - QUESTION: 3
HoE#3V.36.go
Hands-on exercise #3 - Video .36

Using the code from the previous exercise,
1.	At the package level scope, assign the following values to the three variables
a.	for x assign 42
b.	for y assign “James Bond”
c.	for z assign true
2.	in func main
a.	use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
b.	print out the value stored by variable “s”

---------------

N. Level 1 - QUESTION: 4
HoE#4V.37.go
Hands-on exercise #4 - Video .37

●	FYI - nice documentation and new terminology “underlying type”
○	https://golang.org/ref/spec#Types 
For this exercise
1.	Create your own type. Have the underlying type be an int.
2.	create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
3.	in func main
a.	print out the value of the variable “x”
b.	print out the type of the variable “x”
c.	assign 42 to the VARIABLE “x” using the “=” OPERATOR
d.	print out the value of the variable “x”

---------------

N. Level 1 - QUESTION: 5
HoE#5V.38.go
Hands-on exercise #5 - Video .38

Building on the code from the previous example
1.	at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”

  a.	eg:
  type hotdog int

  var x hotdog
  var y int

2.	in func main
  a.	this should already be done
    i.	print out the value of the variable “x”
    ii.	print out the type of the variable “x”
    iii.	assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
    iv.	print out the value of the variable “x”
  b.	now do this
    i.	now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
1.	then use the “=” operator to ASSIGN that value to “y”
2.	print out the value stored in “y”
3.	print out the type of “y”

------------------
------------------

N. Level 2 - QUESTION: 1
HoE#6V.48.go
Hands-on exercise #6 - Video .48

Write a program that prints a number in decimal, binary, and hex

--------------

N. Level 2 - QUESTION: 2
HoE#7V.49.go
Hands-on exercise #7 - Video .49

Using the following operators, write expressions and assign their values to variables:
g.	==
h.	<=
i.	>=
j.	!=
k.	<
l.	>
Now print each of the variables. 

---------------------

N. Level 2 - QUESTION: 3
HoE#8V.50.go
Hands-on exercise #8 - Video .50

Create TYPED and UNTYPED constants. Print the values of the constants.

//Type & untyped constants Explanations:
//Constants may be typed or untyped. 
//Literal constants, true, false, iota, 
//and certain constant expressions 
//containing only untyped constant 
//operands are untyped.

-----------------------

N. Level 2 - QUESTION: 4
HoE#9V.51.go
Hands-on exercise #9 - Video .51

Write a program that 
●	assigns an int to a variable
●	prints that int in decimal, binary, and hex
●	shifts the bits of that int over 1 position to the left, and assigns that to a variable
●	prints that variable in decimal, binary, and hex

--------------------

N. Level 2 - QUESTION: 5
HoE#10V.52.go
Hands-on exercise #10 - Video .52

Create a variable of type string using a raw string literal. Print it.

---------------------

N. Level 2 - QUESTION: 6
HoE#11V.53.go
Hands-on exercise #11 - Video .53

Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

----------------------
----------------------

N. Level 3 - QUESTION: 1
HoE#12V.67.go
Hands-on exercise #12 - Video .67

Print every number from 1 to 10,000

----------------------

N. Level 3 - QUESTION: 2
HoE#13V.68.go
Hands-on exercise #13 - Video .68

Print every rune code point of the uppercase alphabet three times. Your output should look like this:
65
	U+0041 'A'
	U+0041 'A'
  U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B' 
 … through the rest of the alphabet characters

 ------------------

N. Level 3 - QUESTION: 3
HoE#14V.69.go
Hands-on exercise #14 - Video .69

 Create a for loop using this syntax
●	for condition { }
Have it print out the years you have been alive.

-------------------

N. Level 3 - QUESTION: 4
HoE#15V.70.go
Hands-on exercise #15 - Video .70

Create a for loop using this syntax
●	for { }
Have it print out the years you have been alive.

---------------------

N. Level 3 - QUESTION: 5
HoE#16V.71.go
Hands-on exercise #16 - Video .71

Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.

--------------------

N. Level 3 - QUESTION: 6
HoE#17V.72.go
Hands-on exercise #17 - Video .72

Create a program that shows the “if statement” in action.

--------------------

N. Level 3 - QUESTION: 7
HoE#18V.73.go
Hands-on exercise #18 - Video .73

Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

--------------------

N. Level 3 - QUESTION: 8
HoE#19V.74.go
Hands-on exercise #19 - Video .74

Create a program that uses a switch statement with no switch expression specified.

-------------------

N. Level 3 - QUESTION: 9
HoE#20V.75.go
Hands-on exercise #20 - Video .75

Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”

--------------------

N. Level 3 - QUESTION: 10
HoE#21V.76.go
Hands-on exercise #21 - Video .76

Write down what these print:
●	fmt.Println(true && true) 
●	fmt.Println(true && false) 
●	fmt.Println(true || true) 
●	fmt.Println(true || false) 
●	fmt.Println(!true)

-------------------
-------------------

N. Level 4 - QUESTION: 1
HoE#22V.88.go
Hands-on exercise #22 - Video .88

●	Using a COMPOSITE LITERAL: 
    ○	create an ARRAY which holds 5 VALUES of TYPE int
    ○	assign VALUES to each index position. 
●	Range over the array and print the values out. 
●	Using format printing
    ○	print out the TYPE of the array

---------------------

N. Level 4 - QUESTION: 2
HoE#23V.89.go
Hands-on exercise #23 - Video .89

●	Using a COMPOSITE LITERAL: 
  ○	 create a SLICE of TYPE int
  ○	 assign 10 VALUES 
●	Range over the slice and print the values out. 
●	Using format printing
  ○	print out the TYPE of the slice

---------------------

N. Level 4 - QUESTION: 3
HoE#24V.90.go
Hands-on exercise #24 - Video .90

Using the code from the previous example, use SLICING to create the following new slices which are then printed:
●	[42 43 44 45 46]
●	[47 48 49 50 51]
●	[44 45 46 47 48]
●	[43 44 45 46 47]

---------------------

N. Level 4 - QUESTION: 4
HoE#25V.91.go
Hands-on exercise #25 - Video .91

Follow these steps:
●	start with this slice
  ○	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
●	append to that slice this value
  ○	52
●	print out the slice
●	in ONE STATEMENT append to that slice these values
  ○	53
  ○	54
  ○	55
●	print out the slice
●	append to the slice this slice
  ○	y := []int{56, 57, 58, 59, 60}
●	print out the slice

---------------------

N. Level 4 - QUESTION: 5
HoE#26V.92.go
Hands-on exercise #26 - Video .92

To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
●	start with this slice
  ○	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
●	use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
  ○	[42, 43, 44, 48, 49, 50, 51]

NOTE: This "y" is not in the correct solution not sure why. Also he added another Print out frst. He looked tired in the video :)

----------------------

N. Level 4 - QUESTION: 6
HoE#27V.93.go
Hands-on exercise #27 - Video .93

● Create a slice to store the names of all of the states in the United States of America. 

● What is the length of your slice? 

● What is the capacity? 

● Print out all of the values, along with their index position in the slice, without using the range clause. 

Here is a list of the states:

` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`, 

----------------------

N. Level 4 - QUESTION: 7
HoE#28V.94.go
Hands-on exercise #28 - Video .94

Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.

------------------

N. Level 4 - QUESTION: 8
HoE#29V.95.go
Hands-on exercise #29 - Video .95

Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

--------------------

N. Level 4 - QUESTION: 9
HoE#30V.96.go
Hands-on exercise #30 - Video .96

Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

--------------------

N. Level 4 - QUESTION: 10
HoE#31V.97.go
Hands-on exercise #31 - Video .97

Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop

--------------------
--------------------

N. Level 5 - QUESTION: 1
HoE#32V.103.go
Hands-on exercise #32 - Video .103

Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
●	first name
●	last name
●	favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors. 

-------------------

N. Level 5 - QUESTION: 2
HoE#33V.104.go
Hands-on exercise #33 - Video .104

Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

----------------------


