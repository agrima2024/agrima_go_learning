package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is exectuted")
	panic("Oh no. I can do no more. Goodbye.")
	fmt.Println("This is not executed")
}