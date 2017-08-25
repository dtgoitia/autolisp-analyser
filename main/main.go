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

	// Create the array of FileFuncDepen
	var data []alisp.FuncDepen
	var structChunk alisp.FuncDepen
	for _, chunk := range chunksCleaned {
		fmt.Println("  > Chunk:", chunk)

		structChunk = &alisp.FuncDepen{
			FunctionName: filePath,
			Dependencies: []string{"asd", "sdf"},
		}
		data = append(data, structChunk)
	}
	for i, x := range data {
		fmt.Print("\n" + string(i) + ": " + x.FunctionName + "  ")
		fmt.Print(x.Dependencies)
	}
	// e := alisp.FuncDepen{
	// 	FunctionName: "name",
	// 	Dependencies: []string{"ing", "asd"},
	// }
	// fmt.Println("e:", e)
	// fmt.Println("e.FunctionName:", e.FunctionName)
	// fmt.Println("e.Dependencies:", e.Dependencies)
}
