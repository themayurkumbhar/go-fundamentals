package main

import "fmt"

func main() {
	name := "Mayur Kumbahr"
	course := "go-fundamentals"

	fmt.Println("Hi,", name, "welcome to", course)
	updateCourse(&course) // passing the memory address to function here
	fmt.Println("You are currently doing: ", course, "Cource.")

}

// here we are accessing as pointer variable using *string so course makes pointer var
func updateCourse(course *string) string {
	*course = "Getting Started with Go Advance" // here *course dereference the value and udpate the string to new value
	fmt.Println("Updated course to", *course)   // get new updated value
	return *course
}
