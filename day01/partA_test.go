package day01

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`3   4
4   3
2   5
1   3
3   9
3   3
`,
	)
	expected := 11
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
