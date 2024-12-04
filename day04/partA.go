package template

import (
	"bufio"
	"fmt"
	"io"
)

func PrintMap(m [][]byte) {
	for _, r := range m {
		for _, v := range r {
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

var Word = [...]byte{'X', 'M', 'A', 'S'}
var Directions = [...][2]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
	{1, 1},
	{-1, 1},
	{1, -1},
	{-1, -1},
}

func InBounds(m [][]byte, ki, kj int) bool {
	return ki >= 0 && ki < len(m) && kj >= 0 && kj < len(m[0])
}

func isXmas(m [][]byte, i, j, letter_idx int, dir [2]int) bool {
	ki := i + dir[0]*letter_idx
	kj := j + dir[1]*letter_idx

	if !InBounds(m, ki, kj) {
		return false
	}

	return m[ki][kj] == Word[letter_idx]
}

func countXmas(m [][]byte, i, j int) int {
	count := len(Directions)

	for _, dir := range Directions {
		for letter_idx := 1; letter_idx < len(Word); letter_idx++ {
			if !isXmas(m, i, j, letter_idx, dir) {
				count--
				break
			}
		}
	}
	return count
}

func PartASol(m [][]byte) int {
	acc := 0

	for i, row := range m {
		for j, v := range row {
			if v == 'X' {
				acc += countXmas(m, i, j)
			}
		}
	}

	return acc
}

func PartA(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	m := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Bytes()
		m = append(m, append([]byte{}, line...))
	}

	PrintMap(m)
	acc := PartASol(m)

	return acc
}
