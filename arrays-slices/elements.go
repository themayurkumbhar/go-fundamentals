package main

import "fmt"

func main() {

	data := []int{1, 2, 3, 4, 5}
	//slice of slice data
	dataSlice := data[2:5]
	fmt.Println(dataSlice)
	// 0 to 3 data slice
	fmt.Println(data[:3])
	// 3 to end of list slice
	fmt.Println(data[3:])
}
