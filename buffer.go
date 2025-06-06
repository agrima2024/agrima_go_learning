package main

import (
	"fmt"
	"bytes"
)

func main() {
	var buffer bytes.Buffer

	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}

	fmt.Println(buffer.String())
}