package day07

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
	)
	expected := 3749
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
