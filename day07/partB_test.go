package day07

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
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
	expected := 11387
	result := PartB(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestConcat(t *testing.T) {
	expected := 11387
	result := concat(11, 387)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}

	expected = 10
	result = concat(1, 0)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
