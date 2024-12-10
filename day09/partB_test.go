package day09

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	input := strings.NewReader(
		`2333133121414131402`,
	)
	expected := 2858
	result := PartB(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}
