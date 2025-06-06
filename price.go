package main

import (
    "fmt"
)
 
func price(price int) int {
    discount := int(0.1 * float64(price))
	new_price := price - discount
	return new_price
}

func main() {
   fmt.Println(price(23))

}