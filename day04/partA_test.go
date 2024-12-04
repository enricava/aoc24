package template

import (
	"strings"
	"testing"
)

func TestAA(t *testing.T) {
	input := strings.NewReader(
		`....XXMAS.
.SAMXMS...
...S..A...
..A.A.MS.X
XMASAMX.MM
X.....XA.A
S.S.S.S.SS
.A.A.A.A.A
..M.M.M.MM
.X.X.XMASX`,
	)
	expected := 18
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA(t *testing.T) {
	input := strings.NewReader(
		`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
	)
	expected := 18
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_0(t *testing.T) {
	input := strings.NewReader(
		`XMAS`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_1(t *testing.T) {
	input := strings.NewReader(
    `SAMX`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_2(t *testing.T) {
	input := strings.NewReader(
    `X
M
A
S`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_3(t *testing.T) {
	input := strings.NewReader(
    `S
A
M
X`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
