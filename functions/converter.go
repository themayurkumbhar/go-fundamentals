package main

import (
	"fmt"
	"strings"
)

func main() {
	author := "some author"
	course := "go fundamentals"

	fmt.Println(converter(author, course))
}

// function name  params		type		output vars				type
func converter(author, course string) (convertedAuthor, convertedCourse string) {
	author = strings.ToUpper(author)
	course = strings.Title(course)
	return author, course
}
