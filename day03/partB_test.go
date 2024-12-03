package day03

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	input := strings.NewReader(
		`don't()do()xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
	)
	expected := 48
	result := PartB(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
