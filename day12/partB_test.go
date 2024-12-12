package template

import (
	"fmt"
	"strings"
	"testing"
)

func printMap(m [][]byte) {
	for _, r := range m {
		for _, v := range r {
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

func TestB(t *testing.T) {
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
	expected := 1206
	result := PartB(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}

func TestB2(t *testing.T) {
	input := strings.NewReader(
		`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
	)
	expected := 368
	result := PartB(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}

