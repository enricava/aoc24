package day02

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	input := strings.NewReader(
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
	)
	expected := 4
	result := PartB(input).(int)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
