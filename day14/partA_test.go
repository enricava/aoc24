package day14

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`,
	)
	expected := 12
	result := solveA(input, 100, 11, 7)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}

func TestEasy(t *testing.T) {
	input := strings.NewReader(
		`p=2,4 v=2,-3`,
	)

	expected := 0
	result := solveA(input, 5, 11, 7)
	println(expected, result)
	if expected != result {
		t.Fail()
	}
}
