package template

import (
	"strings"
	"testing"
)

func TestB(t *testing.T) {
	input := strings.NewReader(
		`.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
`,
	)
	expected := 9
	result := PartB(input)
  println(expected, result)
	if (expected != result) {
	  t.Fail()
	}
}
