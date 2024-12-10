package day10

import (
	"bufio"
	"io"
	"strconv"
)

func countTrailsB(m [][]int, prev, cur Coord, prevHeight int) int {
	i := cur.x
	j := cur.y

	if !inside(m, i, j) {
		return 0
	}

	height := m[i][j]
	if height-prevHeight != 1 {
		return 0
	}

	if height == 9 {
		return 1
	}

	trails := 0
	for _, dir := range dirs {
		next := add(cur, dir)
		if next != prev {
			trails += countTrailsB(m, cur, next, height)
		}
	}

	return trails
}

func sumTrailheadsB(m [][]int, heads []Coord) int {
	acc := 0

	for _, coord := range heads {
		acc += countTrailsB(m, coord, coord, -1)
	}

	return acc
}

func PartB(file io.Reader) int {
	m := make([][]int, 0)
	heads := make([]Coord, 0)

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		for j, v := range line {
			num, _ := strconv.Atoi(string(v))
			row[j] = num

			if num == 0 {
				heads = append(heads, Coord{i, j})
			}
		}

		m = append(m, row)
		i++
	}

	sol := sumTrailheadsB(m, heads)
	return sol
}
