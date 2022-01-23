package main

import "fmt"

func main() {
	totalWatchedMins := 600
	totalConentMins := 600

	if totalConentMins > totalWatchedMins {
		fmt.Println("You have to watch more mins of course!")
	} else if totalWatchedMins == totalConentMins {
		fmt.Println("You have completed the course!")
	} else {
		fmt.Println("Total watch content must not be greater than total content")
	}

}
