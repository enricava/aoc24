package day06

import (
	"bufio"
	"fmt"
	"io"
)

// next time I should use complex numbers
var dirs = [...][3]int{
	// i,j, bit
	{-1, 0, 1 << 1},
	{0, 1, 1 << 2},
	{1, 0, 1 << 3},
	{0, -1, 1 << 4},
}

func markVisited(m [][]int, i, j, dir_index int) {
	dir_bit := dirs[dir_index][2]
	if m[i][j]&dir_bit == 0 {
		m[i][j] |= dir_bit
	}
}

func visitedWithDir(block int, dir_index int) bool {
	dir_bit := dirs[dir_index][2]
	return block&dir_bit > 0
}

func placeBoulder(m [][]int, i, j int) {
	m[i][j] = -1
}

func isBlocked(m [][]int, i, j int) bool {
	return m[i][j] <= 0
}

func PartBSol(m [][]int, i, j int) int {
	acc := 0
	d := 0

	triedBoulder := make([][]bool, len(m))
	for i := 0; i < len(m); i++ {
		triedBoulder[i] = make([]bool, len(m[0]))
	}

	for ni, nj := next(i, j, d); inside(m, ni, nj); ni, nj = next(i, j, d) {
		if isBlocked(m, ni, nj) {
			d = rotate(d)
			continue
		}

		if !triedBoulder[ni][nj] {
			newMap := deepCopyWithBoulder(m, ni, nj)

			if hasLoop(newMap, i, j, d) {
				acc += 1
				// printMap(newMap)
			}
			triedBoulder[ni][nj] = true
		}

		// advance
		i, j = ni, nj
		markVisited(m, i, j, d)
	}

	// printMap(m)
	return acc
}

func deepCopyWithBoulder(m [][]int, ki, kj int) (out [][]int) {
	out = make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		out[i] = make([]int, len(m[0]))
		for j, v := range m[i] {
			out[i][j] = min(v, 1)
		}
	}
	out[ki][kj] = -1
	return
}

func hasLoop(m [][]int, i, j, d int) bool {
	markVisited(m, i, j, d)

	for ni, nj := next(i, j, d); inside(m, ni, nj); ni, nj = next(i, j, d) {
		if isBlocked(m, ni, nj) {
			d = rotate(d)
			continue
		}

		// advance
		i, j = ni, nj
		if visitedWithDir(m[i][j], d) {
			return true
		}
		markVisited(m, i, j, d)
	}

	return false
}

func PartB(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	m := make([][]int, 0)

	ki := 0
	kj := 0
	i := 0

	for scanner.Scan() {
		line := scanner.Bytes()
		row := make([]int, len(line))

		for j, v := range line {
			if v == '#' {
				row[j] = 0
			} else {
				row[j] = 1
			}

			// find start
			if v == '^' {
				ki = i
				kj = j
			}
		}

		i += 1
		m = append(m, row)
	}

	acc := PartBSol(m, ki, kj)
	return acc
}

// please God forgive me for this function
// i am tired
func printMap(m [][]int) {
	slice := []int{}

	for _, r := range m {
		for _, v := range r {
			if !contains(slice, v) {
				slice = append(slice, v)
			}
			// up, right, down, left
			vv := [4]bool{v&2 != 0, v&4 != 0, v&8 != 0, v&16 != 0}

			if v == 0 {
				fmt.Print("# ")
			} else if v < 0 {
				fmt.Print("🪨")
			} else if vv[0] && vv[1] && vv[2] && vv[3] {
				fmt.Print("󰁁 ")
			} else if vv[0] && vv[1] && vv[2] {
				// up, right, down
				fmt.Print("+ ")
			} else if vv[0] && vv[1] && vv[3] {
				// up, right, left
				fmt.Print("+ ")
			} else if vv[0] && vv[2] && vv[3] {
				// up, down, left
				fmt.Print("+ ")
			} else if vv[1] && vv[2] && vv[3] {
				// right, down, left
				fmt.Print("+ ")
			} else if vv[0] && vv[1] {
				// up, right
				fmt.Print("󰁜 ")
			} else if vv[0] && vv[2] {
				// up, down
				fmt.Print("󰹹 ")
			} else if vv[0] && vv[3] {
				// up, left
				fmt.Print("󰁛 ")
			} else if vv[1] && vv[2] {
				// right, down
				fmt.Print("󰁃 ")
			} else if vv[1] && vv[3] {
				// right, left
				fmt.Print(" ")
			} else if vv[2] && vv[3] {
				// down, left
				fmt.Print("󰁂 ")
			} else if vv[0] {
				// up
				fmt.Print(" ")
			} else if vv[1] {
				// right
				fmt.Print(" ")
			} else if vv[2] {
				// down
				fmt.Print(" ")
			} else if vv[3] {
				// left
				fmt.Print(" ")
			} else {
				// none
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	// fmt.Println("slice:", slice)
	fmt.Println("______________________________")
}

func contains(slice []int, v int) bool {
	for _, x := range slice {
		if x == v {
			return true
		}
	}
	return false
}
