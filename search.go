package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    // Set the test folder path
    test_folder := "./test_folder"

    // Read the directory
    files, err := ioutil.ReadDir(test_folder)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Check each file and print only if it's a regular file
    for _, file := range files {
        if !file.IsDir() {
            fmt.Println(file.Name())
        }
    }
}
