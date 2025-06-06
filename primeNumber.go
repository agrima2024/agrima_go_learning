package main

import (
	"fmt"
)
 
func primeNumber(a int) string {
    var c int = 2
	var b string

	if a == 0{
	b = "Your number is a special number"
	}

	if a == 1{
		b = "Your number is a special number"
	}
	
	for c < a-1{
		c++
		var d = a%c

		if d == 0{
			b = "Your number is not a prime number"
			break
		}

		if d != 0 {
			b = "Your number is a prime number"
		}
	}

	fmt.Println(b)
	return b
}

func main() {
	primeNumber(2)
}