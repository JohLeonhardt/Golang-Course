package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello1")
	fmt.Fprintln(os.Stdout, "Hello2")
	io.WriteString(os.Stdout, "Hello3")
}
