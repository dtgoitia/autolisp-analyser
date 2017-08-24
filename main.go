package main

import (
	"fmt"
	"io/ioutil"
)

// error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := "./testFiles/file1.lsp"

	// Read the file
	dat, err := ioutil.ReadFile(filePath)
	check(err)

	// If successfull print
	fmt.Println(string(dat))

}
