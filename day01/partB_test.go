package day01

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	input := strings.NewReader(
		`3   4
4   3
2   5
1   3
3   9
3   3
`,
	)
	expected := 31
	result := PartB(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
