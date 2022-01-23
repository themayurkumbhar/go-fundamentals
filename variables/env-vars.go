package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Environ()) // reads all the ENV vars present in you system

	fmt.Println("User logged is:", os.Getenv("USERNAME"))
}
