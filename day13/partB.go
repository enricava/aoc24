package day13

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func readButtonsB(file io.Reader) int {
	acc := 0
	scanner := bufio.NewScanner(file)

	button_re := regexp.MustCompile(`^Button [AB]: X\+(\d+), Y\+(\d+)`)
	res_re := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	for scanner.Scan() {
		lineA := scanner.Text()
		scanner.Scan()
		lineB := scanner.Text()
		scanner.Scan()
		lineC := scanner.Text()

		axay := button_re.FindStringSubmatch(lineA)
		ax, _ := strconv.Atoi(axay[1])
		ay, _ := strconv.Atoi(axay[2])

		bxby := button_re.FindStringSubmatch(lineB)
		bx, _ := strconv.Atoi(bxby[1])
		by, _ := strconv.Atoi(bxby[2])

		cxcy := res_re.FindStringSubmatch(lineC)
		cx, _ := strconv.Atoi(cxcy[1])
		cy, _ := strconv.Atoi(cxcy[2])

    cx += 10000000000000
    cy += 10000000000000

		d := ax*by - bx*ay

		if d != 0 {
			a := (cx*by - bx*cy) / d
			b := (ax*cy - cx*ay) / d

			xok := a*ax+b*bx == cx
			yok := a*ay+b*by == cy

			if 0 <= a && 0 <= b && xok && yok {
				acc += 3*a + b
			}
		}

		if !scanner.Scan() {
			break
		}
	}

	return acc
}

func PartB(file io.Reader) int {
	return readButtonsB(file)
}
