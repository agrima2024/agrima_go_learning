package main

import (
	"fmt"
	"errors"
)

func main() {
	err := errors.New("Something went wrong")
	if err != nil {
		fmt.Println(err)
	}
}