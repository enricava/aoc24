package day09

import (
	"strings"
	"testing"
)
func TestA_EnoughPadding(t *testing.T) {
	input := strings.NewReader(
		`253`,
	)
	expected := 9
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA_NotEnoughPadding(t *testing.T) {
	input := strings.NewReader(
		`213`,
	)
	expected := 9
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}

func TestA(t *testing.T) {
	input := strings.NewReader(
		`2333133121414131402`,
	)
	expected := 1928
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
