package day16

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"
)

// cells
const (
	wall int = iota - 1
	empty
	S
	E
	pass
)

var dirs = [...]Coord{
	Coord{1, 0},
	Coord{0, 1},
	Coord{-1, 0},
	Coord{0, -1},
}

type Coord struct {
	i, j int
}

func add(a, b Coord) Coord {
	return Coord{a.i + b.i, a.j + b.j}
}

func sub(a, b Coord) Coord {
	return Coord{a.i - b.i, a.j - b.j}
}

func at[T any](m [][]T, pos Coord) T {
	return m[pos.i][pos.j]
}

func set[T any](m [][]T, pos Coord, new T) {
	m[pos.i][pos.j] = new
}

func inside[T any](m [][]T, c Coord) bool {
	i := c.i
	j := c.j
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[i])
}

func debugCoord(c Coord, m string) {
	log.WithFields(log.Fields{
		"i": c.i,
		"j": c.j,
	}).Debug(m)
}

func readMap(scanner *bufio.Scanner) (m [][]int, start, end Coord) {
	m = make([][]int, 0)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		for j, c := range line {
			var cell int
			switch c {
			case '#':
				cell = wall
			case 'S':
				cell = S
				start = Coord{i, j}
			case 'E':
				cell = E
				end = Coord{i, j}
			}
			row[j] = cell
		}

		m = append(m, row)
		i++
	}

	return
}

func cellToByte(cell int) byte {
	switch cell {
	case wall:
		return '#'
	case empty:
		return '.'
	case S:
		return 'S'
	case E:
		return 'E'
	case pass:
		return 'O'
	}

	log.Warnf("Unexpected cell: %v", cell)
	return '.'
}

func debugMap(m [][]int) {
	var sb strings.Builder

	for _, r := range m {
		for _, v := range r {
			sb.WriteByte(cellToByte(v))
		}
		log.Debug(sb.String())
		sb.Reset()
	}
}

func debugSolutions(m [][]int, s map[Coord]bool) {
	var sb strings.Builder

	for i, r := range m {
		for j, v := range r {
			if _, ok := s[Coord{i, j}]; ok {
				sb.WriteByte('O')
			} else {
				sb.WriteByte(cellToByte(v))
			}
		}
		log.Debug(sb.String())
		sb.Reset()
	}
}

func debugDistance(m [][]int) {
	for _, r := range m {
		for _, v := range r {
			if v == 0 {
				fmt.Print("###### ")
			} else {
				fmt.Printf("%06d ", v)
			}
		}
		fmt.Println()
	}
}

type CoordDir struct {
	coord Coord
	dir   int
}

type Item struct {
	coord    Coord
	dir      int
	distance int
	prev     []Coord
}

func adjs(s Item) []Item {
	visited := append([]Coord{}, s.prev...)
	visited = append(visited, s.coord)

	next := Item{add(s.coord, dirs[s.dir]), s.dir, s.distance + 1, visited}

	nextDir := (s.dir + 1) % 4
	rotateA := Item{add(s.coord, dirs[nextDir]), nextDir, s.distance + 1001, visited}

	nextDir = (s.dir + 3) % 4
	rotateB := Item{add(s.coord, dirs[nextDir]), nextDir, s.distance + 1001, visited}

	return []Item{next, rotateA, rotateB}
}

func dijkstra(s Coord, m [][]int, dist map[CoordDir]int) map[Coord]bool {
	solutions := make(map[Coord]bool)
	q := make([]Item, 0)
	q = append(q, Item{s, 1, 0, []Coord{}})
	dist[CoordDir{s, 1}] = 0
	best := 100000000

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if d, ok := dist[CoordDir{cur.coord, cur.dir}]; ok && cur.distance > d {
			continue
		}

		// visit adjacents
		for _, adj := range adjs(cur) {
			if !inside(m, adj.coord) || at(m, adj.coord) == wall {
				continue
			}

			curBest, seen := dist[CoordDir{adj.coord, adj.dir}]

			if !seen || adj.distance <= curBest {
				dist[CoordDir{adj.coord, adj.dir}] = adj.distance
			}

			if at(m, adj.coord) == E && adj.distance <= best {
				if adj.distance < best {
					best = adj.distance
					solutions = make(map[Coord]bool)
				}
				for _, k := range adj.prev {
					solutions[k] = true
				}
			}

			q = append(q, adj)
		}
	}

	return solutions
}

func PartA(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	m, start, end := readMap(scanner)

	debugCoord(start, "Start")
	debugCoord(end, "End")
	debugMap(m)

	dist := make(map[CoordDir]int)
  dijkstra(start, m, dist)

	sol := 10000000000
	for i, _ := range dirs {
		if d, ok := dist[CoordDir{end, i}]; ok {
			sol = min(sol, d)
		}
	}

	return sol
}
