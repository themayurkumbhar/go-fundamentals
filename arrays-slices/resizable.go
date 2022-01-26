package main

import "fmt"

func main() {
	baseSlice := make([]int, 2, 4)
	fmt.Println("baseslice length is:", len(baseSlice), "with capacity", cap(baseSlice))
	fmt.Println("will start adding elements and see how the size and capacity will grow")

	for i := 0; i < 10; i++ {
		baseSlice = append(baseSlice, i+1)
		fmt.Println("baseslice length is:", len(baseSlice), "with capacity", cap(baseSlice))
	}

	//output will show you capacity is increase by 2x!
}
