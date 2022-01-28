package main

import "fmt"

func main() {
	type titlesMeta struct {
		name   string
		lenght int
		course string
	}

	// var myCourseTitle *titlesMeta
	// myCourseTitle = new(titlesMeta)
	// myCourseTitle.course = "this"
	// myCourseTitle.lenght = 1
	// myCourseTitle.name = "that"
	// fmt.Println(myCourseTitle)

	myCourseTitle := titlesMeta{
		name:   "go-structs-basics",
		lenght: 120,
		course: "go-fundamentals",
	}

	fmt.Println(myCourseTitle)

}
