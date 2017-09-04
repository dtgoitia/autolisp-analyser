package alisp

import (
	"fmt"
	"strings"
)

// FuncDepen : depencency function struct
type FuncDepen struct {
	FunctionName string
	Dependencies []string
}

// FileFuncDepen : depencency file struct
type FileFuncDepen struct {
	FilePath  string
	Functions []FuncDepen
}

// MinifyString : clean unnecessary content from s string
func MinifyString(s string) string {

	patternArray := [][]string{
		{"\r\n", ""}, // newlines (Windows)
		{"\n", ""},   // newlines
		{"  ", ""},   // double spaces
		{" (", "("},  // spaces before "("
		{" )", ")"},  // spaces before ")"
		{"( ", "("},  // spaces after "("
		{") ", ")"},  // spaces after ")"
	}

	newString := s

	for i := range patternArray {
		replacer := strings.NewReplacer(patternArray[i][0], patternArray[i][1])
		newString = replacer.Replace(newString)
	}

	return newString
}

// Chunk : split string in chunks
func Chunk(s string) []string {

	startPoint := 0
	currentNestedLevel := 0
	chunkArray := []string{}

	// Run through every character within "s" string
	for i, x := range s {
		currentChar := string(x)

		switch currentChar {
		case "(":
			if currentNestedLevel == 0 {
				startPoint = i
			}
			currentNestedLevel++

		case ")":
			currentNestedLevel--
			if currentNestedLevel == 0 {
				stringUpToHere := s[startPoint:(i + 1)]
				chunkArray = append(chunkArray, stringUpToHere)
			}
		}

	}
	return chunkArray
}

// CleanChunks : keep only with "DT:*" or "defun" functions
func CleanChunks(chunkArray []string) []string {
	var newChunkArray []string
	for _, x := range chunkArray {
		if x[0:7] == "(defun " {
			newChunkArray = append(newChunkArray, x)
		}
		if x[0:4] == "(DT:" {
			newChunkArray = append(newChunkArray, x)
		}
	}
	return newChunkArray
}

// StringToFileFuncDepen : convert string array FileFuncDepen array
func StringToFileFuncDepen(chunkArray []string, filePath string) FileFuncDepen {
	fmt.Println("Converting strings into structs...")

	var funcDepenArray []FuncDepen
	var fileFuncDepenArray FileFuncDepen
	var tempFuncDepen FuncDepen
	var functionName string
	var stopCopying bool

	for _, x := range chunkArray {
		stopCopying = false // reset stopCopying
		functionName = ""   // reset functionName

		// if chunk starts with "(defun "
		if x[0:7] == "(defun " {
			for _, ch := range x[7:] { // get from 7th character on until you find "(" or " "
				if ch == '(' || ch == ' ' {
					stopCopying = true
				}
				if stopCopying == false {
					functionName = functionName + string(ch)
				}
			}
			tempFuncDepen = FuncDepen{
				FunctionName: functionName,
				Dependencies: []string{"dep1", "dep2"},
			}
			funcDepenArray = append(funcDepenArray, tempFuncDepen)
		} else if x[0:4] == "(DT:" { // get from 4th character on until you find "(" or " "
			functionName = "(DT:"
			for _, ch := range x[4:] {
				if ch == '(' || ch == ' ' {
					stopCopying = true
				}
				if stopCopying == false {
					functionName = functionName + string(ch)
				}
			}
			tempFuncDepen = FuncDepen{
				FunctionName: functionName + ")",
				Dependencies: []string{"this is dep 1", "this is dep dep2"},
			}
			funcDepenArray = append(funcDepenArray, tempFuncDepen)
		}

		fileFuncDepenArray = FileFuncDepen{
			FilePath:  filePath,
			Functions: funcDepenArray,
		}
	}
	return fileFuncDepenArray
}
