package day03

import (
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	input := strings.NewReader(
		`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
	)
	expected := 161
	result := PartA(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
