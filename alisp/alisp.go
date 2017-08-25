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

// MinifyString : clean unnecessary content from s string
func MinifyString(s string) string {
	fmt.Println("Minifying string:")
	replacer1 := strings.NewReplacer(
		"\r\n", "", // remove newlines (in Windows is "\r\n", not "\n")
		"  ", " ", // remove double spaces
	)
	fmt.Println("  > Removing \"\\n\"...")
	fmt.Println("  > Removing double spaces...")
	newString := replacer1.Replace(s)

	replacer2 := strings.NewReplacer(
		" )", ")", // remove spaces before "("
		" (", "(", // remove spaces before ")"
	)
	replacer3 := strings.NewReplacer(
		"( ", "(", // remove spaces after "("
		") ", ")", // remove spaces after ")"
	)

	fmt.Println("  > Removing spaces before and after \"(\" and \")\"...")
	newString = replacer2.Replace(newString)
	newString = replacer3.Replace(newString)

	return newString
}

// Chunk : clean unnecessary content from s string
func Chunk(s string) []string {
	fmt.Println("Chunking string...")

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

		// remainingString := s[i:]
		// remainingStringLength := utf8.RuneCountInString(remainingString)
		// posibleDefunString := ""
		// // fmt.Print("current character (" + strconv.Itoa(i) + ") : " + currentChar + "   remainingStringLength: " + strconv.Itoa(remainingStringLength))
		// if remainingStringLength >= 7 {
		// 	posibleDefunString = remainingString[0:7]
		// 	// if posibleDefunString == "(defun " {
		// 	// 	// fmt.Println("   !!!!  ")
		// 	// 	fmt.Print("current character (" + strconv.Itoa(i) + ") : " + currentChar + "   remainingStringLength: " + strconv.Itoa(remainingStringLength))
		// 	// 	fmt.Print("    posible: " + posibleDefunString)
		// 	// }
		// } else {
		// 	posibleDefunString = ""
		// }
		// fmt.Println("chunkArray:", chunkArray)
	}

	return chunkArray
}

// CleanChunks : clean unnecessary content from s string
func CleanChunks(s string) []string {
	fmt.Println("s")
}
