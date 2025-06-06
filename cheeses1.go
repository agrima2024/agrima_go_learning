package main 

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgonge"
	cheeses = append(cheeses, "Camembert")
	fmt.Println(cheeses[2])
}