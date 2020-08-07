# Golang-Course

QUESTION: 1
Hands-on exercise #1 - Video .34
1.	Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
a.	42
b.	“James Bond”
c.	true
2.	Now print the values stored in those variables using 
a.	a single print statement
b.	multiple print statements

-----------

QUESTION: 2
Hands-on exercise #2 - Video .35
1.	Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
a.	identifier “x” type int
b.	identifier “y” type string
c.	identifier “z” type bool
2.	in func main
a.	print out the values for each identifier
b.	The compiler assigned values to the variables. What are these values called?

-------------

QUESTION: 3
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

QUESTION: 4
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

QUESTION: 5
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

-------------

QUESTION: 6
Hands-on exercise #6 - Video .48
Write a program that prints a number in decimal, binary, and hex

--------------

QUESTION: 7
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

QUESTION: 8
Hands-on exercise #8 - Video .50
Create TYPED and UNTYPED constants. Print the values of the constants.

//Type & untyped constants Explanations:
//Constants may be typed or untyped. 
//Literal constants, true, false, iota, 
//and certain constant expressions 
//containing only untyped constant 
//operands are untyped.
