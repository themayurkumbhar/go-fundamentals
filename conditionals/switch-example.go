package main

import "fmt"

func main() {
	switch "go-fundamentals" {
	case "go-advance":
		fmt.Println("welcome to go-advance")
	case "go-services":
		fmt.Println("welcome to go-services")
	case "go-fundamentals":
		fmt.Println("welcome to go-fundamentals")
		fallthrough
	case "go-falls":
		fmt.Println("go with fallthrough")
	default:
		fmt.Println("Welcome!")
	}
}
