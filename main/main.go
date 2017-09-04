package main

import (
	"autolisp-analyser/alisp"
	"fmt"
	"io/ioutil"
	"strconv"
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
	fmt.Println("Minifying string...")
	fileContentMinified := alisp.MinifyString(fileContent)
	// fmt.Println("\nfileContentMinified:\n" + fileContentMinified)

	// Chunk string to get top levels functions
	fmt.Println("Chunking string...")
	fileContentChunked := alisp.Chunk(fileContentMinified)
	// for _, chunk := range fileContentChunked {
	// 	fmt.Println("  > Chunk:", chunk)
	// }

	// Get rid of functions that are not "DT:*" or "defun" functions
	fmt.Println("Cleaning unnecesary functions...")
	chunksCleaned := alisp.CleanChunks(fileContentChunked)

	// Get a FileFuncDepen struct, which contains the path of the file analysed
	// and an array of FuncDepen strcuts, each of which contain the name of the
	// function and the name of the dependencies
	arrayFileFuncDepen := alisp.StringToFileFuncDepen(chunksCleaned, filePath)
	fmt.Println("\nAnalysing file \"" + arrayFileFuncDepen.FilePath + "\":")
	for i, item := range arrayFileFuncDepen.Functions {
		fmt.Print("Function " + strconv.Itoa(i) + ": ")
		fmt.Println(item.FunctionName)
		for ii, dep := range item.Dependencies {
			fmt.Print("   Dep. " + strconv.Itoa(i) + "." + strconv.Itoa(ii) + ": ")
			fmt.Println(dep)
		}
		// fmt.Print(item.Dependencies, "\t")
		// fmt.Print(item.FunctionName, "\t")
	}
}
