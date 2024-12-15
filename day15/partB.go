package day15

import (
	"bufio"
	"io"

	log "github.com/sirupsen/logrus"
)

func readMapB(scanner *bufio.Scanner) (m [][]int, robot Coord) {
	m = make([][]int, 0)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			// finished map
			break
		}

		row := make([]int, len(line)*2)

		for j, c := range line {
			cell := cellFromRune(c)

			switch cell {
			case bot:
				row[j*2] = bot
				row[j*2+1] = empty
				robot = Coord{i, j * 2}

			case box:
				row[j*2] = boxL
				row[j*2+1] = boxR

			case wall:
				row[j*2] = wall
				row[j*2+1] = wall
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

type CoordAndType struct {
	coord Coord
	cell  int
}

func cascadeMoveBot(robot Coord, move int, m [][]int) (nextRobot Coord) {
	at := func(pos Coord) int { return at(m, pos) }
	set := func(pos Coord, next int) { set(m, pos, next) }

	dir := dirFromMove(move)
	nextRobot = robot

	boxesToMove := make(map[Coord]int)
	boxParts := make(map[Coord]int)

	nextCoord := add(robot, dir)
	nextCell := at(nextCoord)

	if nextCell == wall {
		return
	}

	// process the box, if any
	if nextCell == boxL || nextCell == boxR {
		boxParts[nextCoord] = nextCell
	}

	// while there are boxes left to process
	for len(boxParts) > 0 {
		boxRow := make(map[Coord]int)
		for coord, box := range boxParts {
			boxRow[coord] = box
		}
		// extend the row of boxes to fill the gaps if going up/down
		if move != left && move != right {
			for coord, box := range boxParts {
				if box == boxL {
					// add box to the right
					boxRow[add(coord, Coord{0, 1})] = boxR
				} else if box == boxR {
					// add box to the left
					boxRow[add(coord, Coord{0, -1})] = boxL
				}
			}
		}

		boxParts = make(map[Coord]int)
		// get the next partial row of boxes, which means advance from each box in the row
		for coord, box := range boxRow {
			// add this box to be moved
			boxesToMove[coord] = box

			nextCoord := add(coord, dir)
			nextCell := at(nextCoord)
			if nextCell == boxL || nextCell == boxR {
				boxParts[nextCoord] = nextCell
			} else if nextCell == wall {
				// oops, found a wall in the next cell, cannot move
				return nextRobot
			}
		}
	}

	// clear the boxes positions
	for coord, _ := range boxesToMove {
		set(coord, empty)
	}

	// move the boxes
	for coord, box := range boxesToMove {
		next := add(coord, dir)
		set(next, box)
	}

	// move the bot
	nextRobot = add(robot, dir)
	set(robot, empty)
	set(nextRobot, bot)

	return
}

func simulateB(m [][]int, moves []int, robot Coord) {
	for _, move := range moves {
		robot = cascadeMoveBot(robot, move, m)
		debugMap(m)
	}
}

func PartB(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	m, robot := readMapB(scanner)
	moves := readMoves(scanner)

	debugMap(m)
	simulateB(m, moves, robot)
	debugMap(m)

	sol := getBoxValue(m)
	return sol
}
