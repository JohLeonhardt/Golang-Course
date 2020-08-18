package main

import (
	"fmt"
)

func main() {
	foo()
	func() {
		fmt.Println("Anonymouse func ran")
	}()

}
func foo() {
	fmt.Println("foo ran")
}
