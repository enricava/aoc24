package day02

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
	)
	expected := 2
	result := PartA(input).(int)
  println(expected)
  println(result)
	if (expected != result) {
	  t.Fail()
	}
}
