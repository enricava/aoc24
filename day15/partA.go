package day15

import (
	"bufio"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"
)

// cells
const (
	wall int = iota - 1
	empty
	box
	boxL
	boxR
	bot
)

// moves
const (
	up int = iota
	down
	left
	right
)

type Coord struct {
	i, j int
}

func sub(a, b Coord) Coord {
	return Coord{a.i - b.i, a.j - b.j}
}

func add(a, b Coord) Coord {
	return Coord{a.i + b.i, a.j + b.j}
}

func at[T any](m [][]T, pos Coord) T {
	return m[pos.i][pos.j]
}

func set[T any](m [][]T, pos Coord, new T) {
	m[pos.i][pos.j] = new
}

func inside[T any](m [][]T, pos Coord) bool {
	return pos.i >= 0 && pos.i < len(m) && pos.j >= 0 && pos.j < len(m[pos.i])
}

func cellToByte(cell int) byte {
	switch cell {
	case wall:
		return '#'
	case empty:
		return '.'
	case box:
		return 'O'
	case boxL:
		return '['
	case boxR:
		return ']'
	case bot:
		return '@'
	}

	log.Warnf("Unexpected cell: %v", cell)
	return '.'
}

func cellFromRune(char rune) int {
	switch char {
	case 'O':
		return box
	case '#':
		return wall
	case '@':
		return bot
	case '.':
		return empty
	}

	log.Warnf("Unexpected map cell: %c", char)
	return empty
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

func readMap(scanner *bufio.Scanner) (m [][]int, robot Coord) {
	m = make([][]int, 0)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			// finished map
			break
		}

		row := make([]int, len(line))

		for j, c := range line {
			row[j] = cellFromRune(c)
			if row[j] == bot {
				robot = Coord{i, j}
			}
		}

		m = append(m, row)
		i++
	}

	log.WithFields(log.Fields{
		"i": robot.i,
		"j": robot.j,
	}).Debug("Robot")

	return m, robot
}

func readMoves(scanner *bufio.Scanner) (moves []int) {
	moves = make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		for _, c := range line {
			switch c {
			case '^':
				moves = append(moves, up)
			case 'v':
				moves = append(moves, down)
			case '<':
				moves = append(moves, left)
			case '>':
				moves = append(moves, right)
			}
		}
	}

	log.WithFields(log.Fields{
		"count": len(moves),
	}).Debug("Got moves")

	return moves
}

func dirFromMove(move int) (dir Coord) {
	switch move {
	case up:
		dir = Coord{-1, 0}
	case down:
		dir = Coord{1, 0}
	case left:
		dir = Coord{0, -1}
	case right:
		dir = Coord{0, 1}
	default:
		log.WithFields(log.Fields{"move": move}).Panic("Unexpected move")
	}

	return dir
}

func simpleMoveBot(robot, dir Coord, m [][]int) (nextRobot Coord) {
	at := func(pos Coord) int { return at(m, pos) }
	set := func(pos Coord, next int) { set(m, pos, next) }

	// no movement
	nextRobot = robot

	// find last box in a row in the direction we are heading
	next := add(robot, dir)
	for at(next) == box {
		next = add(next, dir)
	}

	switch at(next) {
	case wall: // cannot move
	case box: // cannot move
	case empty:
		// if the previous cell was a box, move the box
		prev := sub(next, dir)
		if at(prev) == box {
			set(next, box)
		}

		// move the bot
		nextRobot = add(robot, dir)
		set(robot, empty)
		set(nextRobot, bot)
	}

	return nextRobot
}

func simulate(m [][]int, moves []int, robot Coord) {
	for _, move := range moves {
		dir := dirFromMove(move)
		robot = simpleMoveBot(robot, dir, m)
	}
	debugMap(m)
}

func getBoxValue(m [][]int) (acc int) {
	for i, row := range m {
		for j, cell := range row {
			if cell == box {
				acc += 100*i + j
			} else if cell == boxL {
				acc += 100*i + j
			}
		}
	}

	return
}

func PartA(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	m, robot := readMap(scanner)
	moves := readMoves(scanner)

	simulate(m, moves, robot)

	sol := getBoxValue(m)
	return sol
}
