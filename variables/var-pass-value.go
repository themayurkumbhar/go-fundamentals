package main

import (
	"fmt"
)

func main() {
	name := "Mayur Kumbahr"
	course := "go-fundamentals"

	fmt.Println("Hi,", name, "welcome to", course)
	updateCourse(course)
	fmt.Println("You are currently doing: ", course, "Cource.")
}

func updateCourse(course string) string {
	course = "Getting Started with Go Advance"
	fmt.Println("Updated course to", course)
	return course
}
