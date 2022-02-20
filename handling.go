package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, fileErr := ioutil.ReadFile("file.txt")
	if fileErr != nil {
		fmt.Println("Error Reading File: ", fileErr)
	} else {
		fmt.Println(string(file))
	}

}
