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

	// in statement declartion of vars which only available if-else scope

	if userAge, requiredAge := 15, 18; userAge > requiredAge {
		fmt.Println("User is allowed")
	} else {
		fmt.Println("User is not allowed:", userAge)
	}

}
