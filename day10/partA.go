package day10

import (
	"bufio"
	"io"
	"strconv"
)

type Coord struct {
	x, y int
}

var dirs = [...]Coord{
	Coord{1, 0},
	Coord{-1, 0},
	Coord{0, 1},
	Coord{0, -1},
}

func inside[T any](m [][]T, i, j int) bool {
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[i])
}

func add(a, b Coord) Coord {
	return Coord{a.x + b.x, a.y + b.y}
}

func countTrails(m [][]int, prev, cur Coord, prevHeight int, visited map[Coord]bool) int {
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
		if !visited[cur] {
			visited[cur] = true
			return 1
		}
		return 0
	}

	trails := 0
	for _, dir := range dirs {
		next := add(cur, dir)
		if next != prev {
			trails += countTrails(m, cur, next, height, visited)
		}
	}

	return trails
}

func sumTrailheads(m [][]int, heads []Coord) int {
	acc := 0

	for _, coord := range heads {
		visited := make(map[Coord]bool, 0)
		acc += countTrails(m, coord, coord, -1, visited)
	}

	return acc
}

func PartA(file io.Reader) int {
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

	sol := sumTrailheads(m, heads)
	return sol
}
