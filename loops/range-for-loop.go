package main

import "fmt"

func main() {

	// ignore the syntax for now
	courses := []string{"course1", "course2", "course3", "course4"}

	// first is always index, if u want to ignore index use '_' like for _, course := range courses {}
	for index, course := range courses {
		fmt.Println("value at", index, "is", course)
	}
}
