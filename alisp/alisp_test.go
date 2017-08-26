package alisp

import "testing"

// TestMinifyString : MinifyString test 1
func TestMinifyString(t *testing.T) {
	expected := "(this is a clean(output))"
	actual := MinifyString("(  this\n     is a clean  (  output )  )  ")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, actual)
	}
}

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
