package day14

import (
	"fmt"
	"io"
)

func makeVisitedMap(w, h int) (visited [][]bool) {
	visited = make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}
	return visited
}

func printVisited(m [][]bool) {
	for _, r := range m {
		for _, v := range r {
			if v {
				fmt.Print("")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func solveB(file io.Reader, s, w, h int) {
	ppx, ppy, vvx, vvy := readRobots(file)

	for ks := 1; ks < s; ks++ {
		fmt.Println("--------------------", ks, "--------------------")
		v := makeVisitedMap(w, h)
		ok := true
		for i := 0; i < len(ppx); i++ {
			px, py := robotAfter(ppx, ppy, vvx, vvy, i, ks, w, h)
			if v[py][px] {
				ok = false
				break
			}
			v[py][px] = true
		}
		if ok {
			printVisited(v)
		}
	}
}

func PartB(file io.Reader) int {
	solveB(file, 100000, 101, 103)
	return 0
}
