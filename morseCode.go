package main

import (
    "fmt"
	"regexp"
)

func morseCode(word string) string {
	var alphabet = [26]string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var morse = [26]string {".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	
	for i := 0; i<len(word); i++ {
		var b int = 0
		for b = 0; b<len(alphabet); b++ { 
			a, err := regexp.MatchString((alphabet[b]), word)
			if err != nil {
				fmt.Println(nil)
			}
			
			if a == true{
				fmt.Println(morse[b])
			}
		}
	}
	return ""
}

func main() {
	morseCode("nitin")
}