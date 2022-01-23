package main

import "fmt"

func main() {

	noOfClipsWatched := watchedTillNow(1, 2, 3, 5)

	fmt.Println("watched till now: ", noOfClipsWatched)

	noOfClipsWatched = watchedTillNow(1, 2, 3, 4, 5, 6)
	// varing param that above
	fmt.Println("watched till now:", noOfClipsWatched)
}

// func keyword function name, params here "...int" specifies as varying param
func watchedTillNow(clips ...int) int {
	clipsCount := 0
	// ignore the for loop for now
	for _, c := range clips {
		clipsCount++
		c++
	}
	return clipsCount
}
