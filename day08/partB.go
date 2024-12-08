package day08

import (
	"bufio"
	"io"
)

func reflections(a, b Coord, hasAntinode [][]bool) int {
	acc := 0

	if a == b {
		if !hasAntinode[a.x][a.y] {
			hasAntinode[a.x][a.y] = true
      return 1
		}
    return 0
	}

	for c := reflect(a, b); inside(hasAntinode, c.x, c.y); c = reflect(a, b) {
		if !hasAntinode[c.x][c.y] {
			hasAntinode[c.x][c.y] = true
			acc += 1
		}

		b = a
		a = c
	}

	return acc
}

func countAntinodesB(m [][]byte, hasAntinode [][]bool) int {
	acc := 0

	antennas := make(map[byte][]Coord, 0)
	mapAntennas(m, antennas)

	// iterate over frequencies
	for _, frequency := range antennas {
		// for each pair of antennas in the freq
    // could start j:=i+1
		for _, a := range frequency {
			for _, b := range frequency {
				acc += reflections(a, b, hasAntinode)
				acc += reflections(b, a, hasAntinode)
			}
		}
	}

	return acc
}

func PartB(file io.Reader) int {
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

	res := countAntinodesB(m, hasAntinode)

	// printMap(m)
	// printAntinodes(hasAntinode)

	return res
}
