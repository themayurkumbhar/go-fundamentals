package main

import "fmt"

func main() {

	mapInteration := map[int]string{
		1: "A",
		2: "B",
		3: "C",
		4: "D",
		5: "E",
		6: "F",
	}

	for mapKey, mapValue := range mapInteration {
		fmt.Println("Key is:", mapKey, "-> value is:", mapValue)
	}
}
