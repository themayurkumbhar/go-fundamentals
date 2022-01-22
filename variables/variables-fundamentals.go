package main

import "fmt"

//declare var outside any func, string will initialize to empty, int will be 0
var (
	name         string
	course       string
	module, clip int // another way to group and initialize the vars
)

func main() {
	fmt.Println("name and course variable is set to: ", name, "and", course, ".")
	fmt.Println("module and clip variable is set to: ", module, "and", clip, ".")
}
