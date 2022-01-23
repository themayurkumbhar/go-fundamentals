package main

import "fmt"

func main() {

	simpleEval := true

	for simpleEval {
		fmt.Println("simple evaluation")
		break
	}

	//traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
