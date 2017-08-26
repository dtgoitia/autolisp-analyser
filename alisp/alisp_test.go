package alisp

import (
	"testing"
)

// TestMinifyString : MinifyString test
func TestMinifyString(t *testing.T) {
	// <setup code>
	// t.Run("SubtestName", ,TestMinifyStringFunc(input, expected))
	t.Run("remove newlines windows", testMinifyStringFunc("(this\r\n\r\n is \r\na clean\r\n(\r\noutput\r\n))  ", "(this is a clean(output))"))
	t.Run("remove newlines", testMinifyStringFunc("(this \n\nis a clean\n(output\n)\n\n\n)  ", "(this is a clean(output))"))
	t.Run("remove double spaces", testMinifyStringFunc("(this     is   a   clean(output))  ", "(this is a clean(output))"))
	t.Run("remove spaces before parenthesis", testMinifyStringFunc("   (this is a clean     (output   )   )", "(this is a clean(output))"))
	t.Run("remove spaces after parenthesis", testMinifyStringFunc("(    this is a clean(    output)   )   ", "(this is a clean(output))"))
	t.Run("AllCases", testMinifyStringFunc("(  this\n     is \r\na clean\r\n \n (  output )  )  ", "(this is a clean(output))"))
}

// testMinifyStringFunc : MinifyString test function
func testMinifyStringFunc(input string, expected string) func(t *testing.T) {
	return func(t *testing.T) {
		actual := MinifyString(input)
		if actual != expected {
			t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
		}
	}
}

// // TestChunk : Chunk test
// func TestChunk(t *testing.T) {
// 	t.Run("Case1", testChunkFunc("(this is a clean(output))(ready(to be)parsed)", []string{"(this is a clean(output))", "(ready(to be)parsed)"}))
// }

// // testChunkFunc : Chunk test function
// func testChunkFunc(input string, expected []string) func(t *testing.T) {
// 	return func(t *testing.T) {
// 		actual := Chunk(input)
// 		for i, x := range actual {
// 			if x != expected[i] {
// 				t.Errorf("Test failed, expected: '%s', got: '%s'\n\texpected: '%s'\n\tactual: '%s'", expected[i], actual[i], expected, actual)
// 			}
// 		}
// 	}
// }

// TestChunk : Chunk test
func TestChunk(t *testing.T) {
	expected := []string{
		"(this is a clean(output))",
		"(ready(to be)parsed)",
	}
	actual := Chunk("(this is a clean(output))(ready(to be)parsed)")
	for i, x := range actual {
		if x != expected[i] {
			t.Errorf("Test failed, expected: '%s', got: '%s'\n\texpected: '%s'\n\tactual: '%s'", expected[i], actual[i], expected, actual)
		}
	}
}
