package main

import "fmt"

func main() {
	totalWatchedMins := 100
	totalConentMins := 600

	if totalConentMins > totalWatchedMins {
		fmt.Println("You have to watch more mins of course!")
	}

}
