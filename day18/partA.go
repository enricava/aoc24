package day18

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	empty = iota
	visited
	corrupted
	bestPath
)

var dirs = [...]Coord{
	Coord{1, 0},  // down
	Coord{0, 1},  // right
	Coord{-1, 0}, // up
	Coord{0, -1}, // left
}

type Map [][]int

type Coord struct {
	i, j int
}

func (a Coord) add(b Coord) Coord {
	return Coord{a.i + b.i, a.j + b.j}
}

func (a Coord) sub(b Coord) Coord {
	return Coord{a.i - b.i, a.j - b.j}
}

func (m *Map) at(pos Coord) int {
	return (*m)[pos.i][pos.j]
}

func (m *Map) set(pos Coord, new int) {
	(*m)[pos.i][pos.j] = new
}

func (m *Map) log() {
	var sb strings.Builder

	for _, r := range *m {
		sb.Reset()
		for _, v := range r {
			switch v {
			case empty:
				sb.WriteString(". ")
			case visited:
				sb.WriteString(" ")
			case bestPath:
				sb.WriteString("󰵹 ")
			case corrupted:
				sb.WriteString("󰝤 ")
			}
		}
		fmt.Println(sb.String())
	}
}

func (c Coord) inside(m Map) bool {
	i := c.i
	j := c.j
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[i])
}

func (c Coord) log() *log.Entry {
	return log.WithFields(log.Fields{
		"i": c.i,
		"j": c.j,
	})
}

type Rain struct {
	bytes     []Coord
	time      int
	timeLimit int
}

func (o *Rain) RainOnce(time int) (ok bool, ret Coord) {
	if time >= len(o.bytes) || time >= o.timeLimit || o.time >= len(o.bytes) {
		ok = false
		return
	}

	ok = true
	ret = o.bytes[o.time]
	o.time++
	return
}

func (o *Rain) RainAll(m *Map) {
	for i := 0; true; i++ {
		more, nextCorrupted := o.RainOnce(i)

		if !more {
			return
		}
		m.set(nextCorrupted, corrupted)
	}
}

func readBytes(file io.Reader) []Coord {
	scanner := bufio.NewScanner(file)
	bytes := []Coord{}

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")

		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])

		bytes = append(bytes, Coord{y, x})
	}

	return bytes
}

func makeMap(r, c int) (m Map) {
	m = make(Map, r)
	for i, _ := range m {
		m[i] = make([]int, c)
	}

	return
}

func (c Coord) adjs(m Map) (adjs []Coord) {
	adjs = []Coord{}

	for _, dir := range dirs {
		next := c.add(dir)
		if next.inside(m) && m.at(next) == empty {
			adjs = append(adjs, next)
		}
	}

	return
}

type Item struct {
	c       Coord
	parents []Coord
}

func bfs(m Map, s, d Coord) int {
	q := []Item{Item{s, []Coord{}}}
	for time := 0; len(q) > 0; time++ {
		for remaining := len(q); remaining > 0; remaining-- {
			cur := q[0]
			q = q[1:]

			if cur.c == d {
        m.set(cur.c, bestPath)
        for _, parent := range cur.parents {
          m.set(parent, bestPath)
        }

				return time
			}

			for _, next := range cur.c.adjs(m) {
				m.set(next, visited)

				nextParents := append([]Coord{}, cur.parents...)
				nextParents = append(nextParents, cur.c)
				q = append(q, Item{next, nextParents})
			}
		}

	}

	return 0
}

func solveA(file io.Reader, side int, timeLimit int) int {
	start := Coord{0, 0}
	end := Coord{side-1, side-1}

	m := makeMap(side, side)
	rain := Rain{readBytes(file), 0, timeLimit}
	rain.RainAll(&m)

	sol := bfs(m, start, end)
	m.log()

	return sol
}

func PartA(file io.Reader) int {
	side := 71
	timelimit := 1024

	return solveA(file, side, timelimit)
}
