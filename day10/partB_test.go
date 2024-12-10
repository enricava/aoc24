package day10

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	input := strings.NewReader(
		`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
	)
	expected := 81
	result := PartB(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}