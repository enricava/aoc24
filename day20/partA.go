package day20

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	empty = iota
	wall
	S
	E
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
				sb.WriteString(".")
			case wall:
				sb.WriteString("#")
			case S:
				sb.WriteString("S")
			case E:
				sb.WriteString("E")
			}
		}
		fmt.Println(sb.String())
	}
	log.WithFields(
		log.Fields{
			"rows(i)": len((*m)),
			"cols(j)": len((*m)[0]),
		}).Debug("Dimensions")
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

func readMap(file io.Reader) (m Map, start, end Coord) {
	s := bufio.NewScanner(file)

	m = make([][]int, 0)
	i := 0
	for s.Scan() {
		l := s.Text()
		r := make([]int, len(l))
		for j, c := range l {
			switch c {
			case '#':
				r[j] = wall
			case 'S':
				r[j] = S
				start = Coord{i, j}
			case 'E':
				r[j] = E
				end = Coord{i, j}
			}
		}
		m = append(m, r)
		i++
	}

	return
}

type Item struct {
	c               Coord
	remainingCheats int
	cheats          []Coord
	prev            Coord
}

func (o Item) nocheatAdjs(m Map) (adjs []Item) {
	adjs = []Item{}

	for _, dir := range dirs {
		next := o.c.add(dir)
		if next != o.prev && next.inside(m) && m.at(next) != wall {
			adjs = append(adjs, Item{next, o.remainingCheats, o.cheats, o.c})
		}
	}

	return
}

func (o Item) cheatAdjs(m Map, testedCheats map[Coord]bool) (adjs []Item) {
	adjs = []Item{}

	for _, dir := range dirs {
		next := o.c.add(dir)
		_, tested := testedCheats[next]
		if !tested && next != o.prev && next.inside(m) && m.at(next) == wall && o.remainingCheats > 0 {
			testedCheats[next] = true
			for t := 0; t < o.remainingCheats; t++ {
				newCheats := append(o.cheats, next)
				adjs = append(adjs, Item{next, t, newCheats, o.c})
			}
		}
	}

	return
}

func bfs(m Map, s, d Coord) int {
	q := []Item{{s, 2, []Coord{}, Coord{}}}
	for time := 0; len(q) > 0; time++ {
		for remaining := len(q); remaining > 0; remaining-- {
			cur := q[0]
			q = q[1:]

			if cur.c == d {
				return time
			}

			for _, next := range cur.nocheatAdjs(m) {
				q = append(q, next)
			}
		}
	}

	return 100000
}

type cheatComb []Coord

func bfsCheat(m Map, s, d Coord, fairScore int, threshold int) int {
	testedCheats := make(map[Coord]bool)
	count := 0
	q := []Item{{s, 1, []Coord{}, Coord{}}}

	for time := 0; len(q) > 0 && time <= fairScore-threshold; time++ {
		for remaining := len(q); remaining > 0; remaining-- {
			cur := q[0]
			q = q[1:]

			if cur.c == d {
				if time <= 20 {
					for _, cc := range cur.cheats {
						cc.log().Debug("Cheated")
					}
				}
				log.Debugf("Cheat scored: %v", time)
				count++
			}

			for _, next := range cur.nocheatAdjs(m) {
				q = append(q, next)
			}

			for _, cheatNext := range cur.cheatAdjs(m, testedCheats) {
				q = append(q, cheatNext)
			}
		}
	}

	return count
}

func PartA(file io.Reader) int {
	m, s, d := readMap(file)

	m.log()
	s.log().Debug("Start")
	d.log().Debug("End")

	noCheat := bfs(m, s, d)
	betterWithCheats := bfsCheat(m, s, d, noCheat, 100)

	log.Debugf("No cheat score: %v", noCheat)

	sol := betterWithCheats

	return sol
}
