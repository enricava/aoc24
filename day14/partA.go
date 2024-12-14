package day14

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func readRobots(file io.Reader) (ppx, ppy, vvx, vvy []int) {
	ppx = make([]int, 0)
	ppy = make([]int, 0)
	vvx = make([]int, 0)
	vvy = make([]int, 0)

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		robot := re.FindStringSubmatch(line)

		px, _ := strconv.Atoi(robot[1])
		py, _ := strconv.Atoi(robot[2])
		vx, _ := strconv.Atoi(robot[3])
		vy, _ := strconv.Atoi(robot[4])

		ppx = append(ppx, px)
		ppy = append(ppy, py)
		vvx = append(vvx, vx)
		vvy = append(vvy, vy)
	}

	return
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func robotAfter(ppx, ppy, vvx, vvy []int, i, s, w, h int) (int, int) {
	px := mod((ppx[i] + (vvx[i] * s)), w)
	py := mod((ppy[i] + (vvy[i] * s)), h)
	return px, py
}

func solveA(file io.Reader, s, w, h int) int {
	ppx, ppy, vvx, vvy := readRobots(file)
	q0, q1, q2, q3 := 0, 0, 0, 0

	for i := 0; i < len(ppx); i++ {
		px, py := robotAfter(ppx, ppy, vvx, vvy, i, s, w, h)
		// fmt.Println("Robot:", px, py)

		if px < w/2 && py < h/2 {
			q0++
		} else if px > w/2 && py < h/2 {
			q1++
		} else if px < w/2 && py > h/2 {
			q2++
		} else if px > w/2 && py > h/2 {
			q3++
		}
	}

	return q0 * q1 * q2 * q3
}

func PartA(file io.Reader) int {
	s := 100
	w, h := 101, 103

	return solveA(file, s, w, h)
}
