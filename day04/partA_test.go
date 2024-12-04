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

func TestA_4(t *testing.T) {
	input := strings.NewReader(
    `S000
0A00
00M0
000X`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_5(t *testing.T) {
	input := strings.NewReader(
    `X000
0M00
00A0
000S`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_6(t *testing.T) {
	input := strings.NewReader(
    `000X
00M0
0AA0
S00S`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_7(t *testing.T) {
	input := strings.NewReader(
    `000S
00A0
0MA0
X00S`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_8(t *testing.T) {
	input := strings.NewReader(
    `0SSS
AAA0
0MM0
XXXS`,
	)
	expected := 3
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_9(t *testing.T) {
	input := strings.NewReader(
    `S0S0S
0AAA0
00M00
00X00`,
	)
	expected := 1
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

