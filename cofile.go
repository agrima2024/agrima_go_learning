package main

import (
	"log"
	"os"
)

func main() {
	from, err := so.Open("./example02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile("./example02.copy.txt", os.O_RDWR|os.O_Create, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close(to, from)

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}