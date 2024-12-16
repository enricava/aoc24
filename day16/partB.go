package day16

import (
	"bufio"
	"io"

	// log "github.com/sirupsen/logrus"
)

func keys(m map[Coord]bool) []Coord {
	i := 0
	keys := make([]Coord, len(m))
	for k, _ := range m {
		keys[i] = k
		i++
	}
	return keys
}

func PartB(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	m, start, end := readMap(scanner)

	debugCoord(start, "Start")
	debugCoord(end, "End")
	debugMap(m)

	dist := make(map[CoordDir]int)
  solutions := dijkstra(start, m, dist)

  debugSolutions(m, solutions)

	return len(solutions) + 1
}
