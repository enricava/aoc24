package day11

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`125 17`,
	)
	expected := 55312
	result := PartA(input)

	println(expected, result)
	if expected != result {
		t.Fail()
	}
}
