package template

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
	)
	expected := 1930
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
