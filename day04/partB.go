package template

import (
	"bufio"
	"io"
)

// a.b
// .A.
// d.c
var Valid = [...][4]byte{
	//a ,  b ,  c ,  d
	{'M', 'M', 'S', 'S'},
	{'M', 'S', 'S', 'M'},
	{'S', 'S', 'M', 'M'},
	{'S', 'M', 'M', 'S'},
}

func isCrossmas(m [][]byte, i, j int) bool {
	ai := i - 1
	aj := j - 1
	if !InBounds(m, ai, aj) {
		return false
	}

	bi := i - 1
	bj := j + 1
	if !InBounds(m, bi, bj) {
		return false
	}

	ci := i + 1
	cj := j + 1
	if !InBounds(m, ci, cj) {
		return false
	}

	di := i + 1
	dj := j - 1
	if !InBounds(m, di, dj) {
		return false
	}

	instance := [4]byte{m[ai][aj], m[bi][bj], m[ci][cj], m[di][dj]}

	for _, valid := range Valid {
		if instance == valid {
			return true
		}
	}

	return false
}

func PartBSol(m [][]byte) int {
	acc := 0

	for i, row := range m {
		for j, v := range row {
			if v == 'A' {
				if isCrossmas(m, i, j) {
					acc += 1
				}
			}
		}
	}

	return acc
}

func PartB(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	m := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		m = append(m, append([]byte{}, line...))
	}

	acc := PartBSol(m)

	return acc
}
