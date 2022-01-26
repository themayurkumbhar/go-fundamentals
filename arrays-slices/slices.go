package main

import (
	"fmt"
)

func main() {

	//initialzing empty slices
	//make is inbuild function to initialize the slices
	// var			type, size, capacity(max size) <- array of this size is crated at backend
	titles := make([]string, 3, 5)
	titles[0] = "arryas and slices"
	titles[2] = "functions"
	titles[1] = "fundamentals of go"
	fmt.Println("Number of titles are:", len(titles), "Capacity can be:", cap(titles))
	fmt.Println(titles)

	authors := []string{
		"Author1",
		"Author4",
		"Author2",
		"Author3"}

	fmt.Println("Authors slice size is:", len(authors), "Capacity is:", cap(authors))

	for _, author := range authors {
		fmt.Println(author)
	}
}
