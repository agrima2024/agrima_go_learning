package main

import (
	"fmt"
)

func reverse_a_word(str string) string {
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse
}

func main() {
	fmt.Println(reverse_a_word("hara"))
}

