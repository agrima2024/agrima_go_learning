package main

import (
	"fmt"
)
func small_and_big() {
	var anything = []int {5,2,90,56}
	var small int
	var big int 
	var i int

	small = anything[0]
	big = anything[0]

	for i = 0; i < len(anything); i++ {
		if small > anything[i] {
			small = anything[i]
		}

		if big < anything[i] {
			big = anything[i]
		}
	}
	fmt.Println("The smallest number is", small)
	fmt.Println("The biggest number is", big)
}


func main() {
	small_and_big()
}