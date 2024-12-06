package day06

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
	)
	expected := 41
	result := PartA(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}
