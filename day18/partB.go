package day18

import (
	"io"

	log "github.com/sirupsen/logrus"
)

type Corruptible struct {
	c         Coord
	parents   map[Coord]bool
	corrupted bool
}

func solveB(file io.Reader, side int, timeLimit int, timestart int) int {
	start := Coord{0, 0}
	end := Coord{side - 1, side - 1}

	m := makeMap(side, side)
	rain := Rain{readBytes(file), 0, timeLimit}

	time := 0

	for time = 0; time < timestart-1; time++ {
		_, corrupt := rain.RainOnce(time)
		m.set(corrupt, corrupted)
	}

	corrupt := Coord{}

	for time = timestart - 1; time < len(rain.bytes); time++ {
		_, corrupt = rain.RainOnce(time)
		m.set(corrupt, corrupted)

		sol := bfs(m, start, end)
		if sol == 0 {
			break
		}
		log.WithFields(log.Fields{"time": time, "sol": sol}).Debug("Has solution")

		// reset the matrix
		for i, r := range m {
			for j, v := range r {
				if v != corrupted {
					m[i][j] = empty
				}
			}
		}
	}

	log.WithFields(log.Fields{
		"time": time,
		"last.x": corrupt.j,
		"last.y": corrupt.i,
	}).Warn("No solution")
	m.log()

	return time
}

func PartB(file io.Reader) int {
	side := 71
	timelimit := 1000000
	timestart := 1024

	return solveB(file, side, timelimit, timestart)
}
