package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutine\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	go bar()
	woo()

	fmt.Println("Goroutine\t", runtime.NumGoroutine())
	wg.Wait()
}
func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
	}
}
func bar() {
	for i := 0; i < 3; i++ {
		fmt.Println("bar:", i)
		wg.Done()
	}
}
func woo() {
	for i := 0; i < 4; i++ {
		fmt.Println("woo:", i)
	}
}
