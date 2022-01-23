package main

import (
	"fmt"
	"reflect"
)

//declare var outside any func, string will initialize to empty, int will be 0
var (
	name, course = "Mayur Kumbhar", "Go-Fundamentals"
	module, clip = 2, 1 // another way to group and initialize the vars
)

func main() {
	fmt.Println("name and course variable is set to: ", name, "and", course, ".")
	fmt.Println("module and clip variable is set to: ", module, "and", clip, ".")
	// chek the type of vars using reflection
	fmt.Println("name is type of: ", reflect.TypeOf(name))
	fmt.Println("module is type of: ", reflect.TypeOf(module))
	var description string // without intializing
	fmt.Println("Description is: ", description)
	title := "Variable Declarations" // with initializing the value to variable
	fmt.Println("title is set: ", title)
	fmt.Println("length is: ", len(title))               // this is pass by value to len funtion
	fmt.Println("Memory address of title var: ", &title) // & sign is used to get memory address

	var ptr *string = &title
	fmt.Print("ptr var pointing to: ", ptr, "memory. which contains value as: ", *ptr)
}
