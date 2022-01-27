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

	//access specific map value
	fmt.Println(mapInteration[5])

	//update specific map value, //if the key dosent exists it will create map entry
	fmt.Println("before update:", mapInteration[3])
	mapInteration[3] = "SomeNewVal"
	fmt.Println("after update:", mapInteration[3])

	//delete entry from map
	fmt.Println("before deletion:", mapInteration)
	delete(mapInteration, 3)
	fmt.Println("after deletion:", mapInteration)
}
