package main

import "fmt"

func main() {

	// ignore the syntax for now
	courses := []string{"course1", "course2", "course3", "course4"}

	// first is always index, if u want to ignore index use '_' like for _, course := range courses {}
	for index, course := range courses {
		fmt.Println("value at", index, "is", course)
	}

	titles := []string{"title1", "title2", "title3", "title4"}

	//nested for loops
	for _, course := range courses {
		for _, title := range titles {
			fmt.Println("Course", course, "has title:", title)
		}
	}
}
