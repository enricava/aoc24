package template

import (
	"bufio"
	"io"
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

func readMap(file io.Reader) (m [][]byte) {
	m = make([][]byte, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, append([]byte{}, line...))
	}

	return m
}

func makeVisitedMap(m [][]byte) (visited [][]bool) {
	visited = make([][]bool, len(m))
	for i, row := range m {
		visited[i] = make([]bool, len(row))
	}
	return visited
}

func visit(m [][]byte, visited [][]bool, i, j int) (int, int) {
	visited[i][j] = true

	area := 1
	fences := 0

	for _, dir := range dirs {
		ki, kj := i+dir.x, j+dir.y

		if inside(m, ki, kj) && !visited[ki][kj] && m[ki][kj] == m[i][j] {
			// same type, no added fences but visit
			a, f := visit(m, visited, ki, kj)
			area += a
			fences += f
		} else if !inside(m, ki, kj) || m[ki][kj] != m[i][j]{
			// different type or out of bounds, so add my fence
			fences++
			// if inside(m, ki, kj) {
			// 	fmt.Printf("at %c[%v, %v] - fence %c[%v, %v]\n", m[i][j], i, j, m[ki][kj], ki, kj)
			// } else {
			// 	fmt.Printf("at %c[%v, %v] - fence -[%v, %v]\n", m[i][j], i, j, ki, kj)
			// }
		}
	}

	return area, fences
}

func settle(m [][]byte) int {
	visited := makeVisitedMap(m)

	acc := 0

	for i, row := range m {
		for j, _ := range row {
			if !visited[i][j] {
				area, fences := visit(m, visited, i, j)
				acc += area * fences
				// fmt.Printf("Visited %c: area %v * fences %v = price %v | total %v\n", v, area, fences, area*fences, acc)
			}
		}
	}

	return acc
}

func PartA(file io.Reader) int {
	m := readMap(file)
	return settle(m)
}
