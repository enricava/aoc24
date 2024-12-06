package day06

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

var Directions = [...][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func next(i, j, d int) (int, int) {
	dir := Directions[d]
	return i + dir[0], j + dir[1]
}

func rotate(d int) int {
	return (d + 1) % 4
}

func inside[T any](m [][]T, i, j int) bool {
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[i])
}

func PartASol(m [][]byte, ki, kj int) int {
	acc := 0

	d := 0
	i := ki
	j := kj

	for ni, nj := next(i, j, d); inside(m, ni, nj); ni, nj = next(i, j, d) {
		if m[i][j] != 'X' {
			m[i][j] = 'X'
			acc++
		}

		if m[ni][nj] == '#' {
			d = rotate(d)
			continue
		}

		i = ni
		j = nj
	}

	if m[i][j] != 'X' {
		m[i][j] = 'X'
		acc++
	}

	return acc
}

func PartA(file io.Reader) int {

	scanner := bufio.NewScanner(file)
	m := make([][]byte, 0)

	ki := 0
	kj := 0

	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		m = append(m, append([]byte{}, line...))

		// find guard
		for j, v := range line {
			if v == '^' {
				ki = i
				kj = j
			}
		}
		i += 1
	}

	acc := PartASol(m, ki, kj)

	return acc
}
