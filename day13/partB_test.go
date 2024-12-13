package day13

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	input := strings.NewReader(
		`myrawstring`,
	)
	expected := 1
	result := PartB(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
