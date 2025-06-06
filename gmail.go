package main

import (
    "fmt"
    "regexp"
)

func email_detecter(str string) string {
    email := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-]`)

    if email.MatchString(str) {
        return "Your email is valid!"
    } else {
        return "Your email is invalid! Try again with a different email."
    }
}

func main() {
    fmt.Println(email_detecter("kishore@startree.ai"))
}