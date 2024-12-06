package day06

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
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
	expected := 6
	result := PartB(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}

func TestB_1(t *testing.T) {
	input := strings.NewReader(
    `####.
....#
....#
...^.
...#.`,
	)
	expected := 3
	result := PartB(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}

func TestB_2(t *testing.T) {
	input := strings.NewReader(
    `.....
#.##.
#..#.
....#
#^...
..##.`,
	)
	expected := 2
	result := PartB(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}

func TestBDesperation(t *testing.T) {
	input := strings.NewReader(
		`##..#.....
.....##..#
...#....#.
..#.....#.
.......#..
..........
....^.....
....#...#.
#.........
......#...`,
	)
	expected := 6
	result := PartB(input)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}
