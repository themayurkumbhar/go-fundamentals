package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// as the main is also go-routine in go, if it complete all the go-routines will complete.
	// so will this to wait for the these to complete
	var waitGrp sync.WaitGroup
	waitGrp.Add(2) //wait for 2 roitines to complete before exit

	// well if you thinking what is code, its actually anonymous functions - a function has no name :P
	// you can create anonymous function just declare function without name
	// add () end of the function which executues the anonymous function
	// adding go infront of function makes it go routine
	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("go-fundamentals")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("learning-concurrency")
	}()

	waitGrp.Wait()
}
