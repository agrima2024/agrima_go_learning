package main

import (
	"fmt"
)

func sequence(limit int) int {
	var a int = 0
	var b int = 1

	for a < limit{
		fmt.Println(a) 
		a = a + b
		b = a - b
	}
	return a 
}

func main() {
	fmt.Println(sequence(1000))
}