package day05

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func skipOne(slice []int, skip_idx int) (newSlice []int) {
	newSlice = make([]int, 0)
	newSlice = append(newSlice, slice[:skip_idx]...)
	newSlice = append(newSlice, slice[skip_idx+1:]...)
	return
}

func moveValue(slice []int, fromIndex, toIndex int) []int {
	value := slice[fromIndex]

	slice = skipOne(slice, fromIndex)

	slice = append(slice[:toIndex], append([]int{value}, slice[toIndex:]...)...)

	return slice
}

func PartB(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	after := make(map[int][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		left, right, ok := strings.Cut(line, "|")

		if !ok {
			break
		}

		num1, _ := strconv.Atoi(left)
		num2, _ := strconv.Atoi(right)

		after[num1] = append(after[num1], num2)
	}

	acc := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Consume numbers
		isIncorrect := false
		numbers := strings.Split(line, ",")
		ordering := []int{}

		for _, numstr := range numbers {
			num, _ := strconv.Atoi(numstr)
			ordering = append(ordering, num)
		}

		for i := 0; i < len(ordering); i++ {
			newpos := i

			// find first issue
			for _, shouldNotHaveSeen := range after[ordering[i]] {
				for j, v := range ordering {
          if j >= i {
            break
          }
					if v == shouldNotHaveSeen {
						isIncorrect = true
						newpos = min(newpos, j)
						break
					}
				}
			}

			if newpos != i {
				ordering = moveValue(ordering, i, newpos)
				// have to recheck from the new position
				// please no cycles :(
				i = newpos + 1
			}
		}

		if isIncorrect {
			for i, v := range ordering {
				if i == len(ordering)/2 {
					acc += v
					// fmt.Print("+")
				}
				// fmt.Print(v, ",")
			}
			// fmt.Println()
		}
	}

	return acc
}
