package main

import (
	"fmt"
	"time"
)

func main() {

	// well if you thinking what is code, its actually anonymous functions - a function has no name :P
	// you can create anonymous function just declare function without name
	// add () end of the function which executues the anonymous function
	func() {
		time.Sleep(5 * time.Second)
		fmt.Println("go-fundamentals")
	}()

	func() {
		fmt.Println("leaning-concurrency")
	}()
}
