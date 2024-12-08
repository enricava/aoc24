package day08

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	input := strings.NewReader(
		`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
	)
	expected := 34
	result := PartB(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
