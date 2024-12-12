package template

import (
	"io"
	"sort"
)

// returns area and modifies vertical / horizontal
func visitDiscount(m [][]byte, visited [][]bool, i, j int, up, down, left, right map[int][]int) int {
	visited[i][j] = true

	area := 1

	for _, dir := range dirs {
		ki, kj := i+dir.x, j+dir.y

		if inside(m, ki, kj) && !visited[ki][kj] && m[ki][kj] == m[i][j] {
			// same type, no added fences but visit
			area += visitDiscount(m, visited, ki, kj, up, down, left, right)
		} else if !inside(m, ki, kj) || m[ki][kj] != m[i][j] {
			// different type or out of bounds, so add my fence
			if i != ki {
				// adds a horizontal fence in the would-go position
				if i > ki {
					left[ki] = append(left[ki], j)
				} else {
					right[ki] = append(right[ki], j)
				}
			} else {
				if j > kj {
					up[kj] = append(up[kj], i)
				} else {
					down[kj] = append(down[kj], i)
				}
			}
		}
	}

	return area
}

func countSides(direction map[int][]int) int {
	sides := 0

	for _, positions := range direction {
		// sort
		sort.Slice(positions, func(i, j int) bool { return positions[i] < positions[j] })

		sides++
		for i := 1; i < len(positions); i++ {
			if positions[i]-positions[i-1] > 1 {
				sides++
			}
		}
	}

	return sides
}

func settleDiscount(m [][]byte) int {
	visited := makeVisitedMap(m)

	acc := 0

	for i, row := range m {
		for j, _ := range row {
			if !visited[i][j] {
				up := make(map[int][]int)
				down := make(map[int][]int)
				left := make(map[int][]int)
				right := make(map[int][]int)

				area := visitDiscount(m, visited, i, j, up, down, left, right)
				sides := countSides(up) + countSides(down) + countSides(left) + countSides(right)
				acc += area * sides
			}
		}
	}

	return acc
}

func PartB(file io.Reader) int {
	m := readMap(file)
	return settleDiscount(m)
}
