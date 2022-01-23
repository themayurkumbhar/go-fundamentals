package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("../README1.md")
	if err != nil {
		fmt.Println("Error opening the file:", err)
	}
}
