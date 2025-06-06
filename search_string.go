package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
)
 
func search_string(folder string) {
	for _, arg := range os.Args {
		files, err := ioutil.ReadDir(folder)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, file := range files {
        if !file.IsDir() {
		data, err := ioutil.ReadFile(folder + "/" + file.Name())
  		if err != nil {
    		fmt.Println("File reading error", err)
    		return
 		}
		 match, err := regexp.MatchString(arg, string(data))
		 if match == true {
			 fmt.Println(arg, "is part of the file", file.Name())
		 }
		 if match == false {
			 fmt.Println(arg, "is not part of the file", file.Name())
		 }
        }
    }
	}
}

func main() {
	search_string("./test_folder")
}