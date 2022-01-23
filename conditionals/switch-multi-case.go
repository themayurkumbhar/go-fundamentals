package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	switch randomNumber := random(); randomNumber {
	case 0, 1, 2, 3:
		fmt.Println("number is in range of 0-3")
	case 4, 6, 7, 8:
		fmt.Println("number is in range of 4-8")
	default:
		fmt.Println("generated:", randomNumber)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
