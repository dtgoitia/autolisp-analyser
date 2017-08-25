package main

import (
	"autolisp-analyser/alisp"
	"fmt"
	"io/ioutil"
)

// check : error handling function
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("\nRUNNING MAIN --------------------------------------------------------------")
	filePath := "./../testFiles/file1.lsp"

	// Read the file
	fileData, err := ioutil.ReadFile(filePath)
	check(err)

	// Convert file data to a string
	fileContent := string(fileData)
	// fmt.Println("\fileContent:\n" + fileContent)

	// Minify file string
	fileContentMinified := alisp.MinifyString(fileContent)
	// fmt.Println("\nfileContentMinified:\n" + fileContentMinified)

	// Chunk string to get top levels functions
	fileContentChunked := alisp.Chunk(fileContentMinified)
	for _, chunk := range fileContentChunked {
		fmt.Println("  > Chunk:", chunk)
	}

	// Get rid of functions that are not "DT:*" or "defun" functions
	chunksCleaned := alisp.CleanChunks(fileContentChunked)
	for _, chunk := range chunksCleaned {
		fmt.Println("  > Chunk:", chunk)
	}
	e := alisp.FuncDepen{
		FunctionName: "name",
		Dependencies: []string{"ing", "asd"},
	}
	fmt.Println("e:", e)
	fmt.Println("e.FunctionName:", e.FunctionName)
	fmt.Println("e.Dependencies:", e.Dependencies)

}
