package day08

import (
	"bufio"
	"fmt"
	"io"
)

type Coord struct {
	x, y int
}

func sub(a, b Coord) Coord {
	return Coord{a.x - b.x, a.y - b.y}
}

func add(a, b Coord) Coord {
	return Coord{a.x + b.x, a.y + b.y}
}

func printMap(m [][]byte) {
	for _, r := range m {
		for _, v := range r {
			fmt.Printf("%c", v)
		}
		fmt.Println()
	}
}

func printAntinodes(m [][]bool) {
	for _, r := range m {
		for _, v := range r {
			if v {
				fmt.Printf("#")
			}
			fmt.Printf(".")
		}
		fmt.Println()
	}
}

func mapAntennas(m [][]byte, antennas map[byte][]Coord) {
	for i, row := range m {
		for j, v := range row {
			if v != '.' {
				antennas[v] = append(antennas[v], Coord{i, j})
			}
		}
	}
}

func inside[T any](m [][]T, i, j int) bool {
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[i])
}

func reflect(a, b Coord) Coord {
	return sub(add(a, a), b)
}

func reflectionPlacesNewAntinode(a, b Coord, hasAntinode [][]bool) bool {
	ref := reflect(a, b)
	if inside(hasAntinode, ref.x, ref.y) && !hasAntinode[ref.x][ref.y] {
		hasAntinode[ref.x][ref.y] = true
		return true
	}
	return false
}

func countAntinodes(m [][]byte, hasAntinode [][]bool) int {
	acc := 0

	antennas := make(map[byte][]Coord, 0)
	mapAntennas(m, antennas)

	// iterate over frequencies
	for _, frequency := range antennas {
		// for each pair of antennas in the freq
    // could start j:=i+1
		for i, a := range frequency {
			for j, b := range frequency {
				if i == j {
					break
				}

				if reflectionPlacesNewAntinode(a, b, hasAntinode) {
					acc += 1
				}

				if reflectionPlacesNewAntinode(b, a, hasAntinode) {
					acc += 1
				}
			}
		}
	}

	return acc
}

func PartA(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	m := make([][]byte, 0)
	hasAntinode := make([][]bool, 0)

	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		m = append(m, append([]byte{}, line...))
		hasAntinode = append(hasAntinode, make([]bool, len(line)))

		i++
	}

	res := countAntinodes(m, hasAntinode)

	// printMap(m)
	// printAntinodes(hasAntinode)

	return res
}
