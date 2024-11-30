package template

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`myrawstring`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
