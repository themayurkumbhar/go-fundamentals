package main

import "fmt"

func main() {
	//create empty map
	// with key type string and value type int
	courses := make(map[string]int)
	//puting values in map.
	courses["go-fundamentals"] = 120
	courses["go-advance"] = 220
	courses["go-oops"] = 190

	fmt.Println("Courses map:", courses)

	authors := map[string]int{
		"Author_One": 1,
		"Author_Two": 2, //this is bcoz go will add ; at end
	}

	fmt.Println("authors are:", authors)

}
