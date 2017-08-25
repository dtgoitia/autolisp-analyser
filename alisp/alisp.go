package alisp

import (
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
	Functions FuncDepen
}

// MinifyString : clean unnecessary content from s string
func MinifyString(s string) string {
	// fmt.Println("Minifying string:")
	replacer1 := strings.NewReplacer(
		"\r\n", "", // remove newlines (in Windows is "\r\n", not "\n")
		"  ", " ", // remove double spaces
	)
	// fmt.Println("  > Removing \"\\n\"...")
	// fmt.Println("  > Removing double spaces...")
	newString := replacer1.Replace(s)

	replacer2 := strings.NewReplacer(
		" )", ")", // remove spaces before "("
		" (", "(", // remove spaces before ")"
	)
	replacer3 := strings.NewReplacer(
		"( ", "(", // remove spaces after "("
		") ", ")", // remove spaces after ")"
	)

	// fmt.Println("  > Removing spaces before and after \"(\" and \")\"...")
	newString = replacer2.Replace(newString)
	newString = replacer3.Replace(newString)

	return newString
}

// Chunk : clean unnecessary content from s string
func Chunk(s string) []string {
	// Run through every character within "s" string
	startPoint := 0
	currentNestedLevel := 0
	chunkArray := []string{}
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
				// fmt.Println("stringUpToHere:", stringUpToHere)
				chunkArray = append(chunkArray, stringUpToHere)
				// fmt.Println("chunkArray:", chunkArray, "\n")
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
