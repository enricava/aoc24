package day05

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func PartA(file io.Reader) int {
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
		numbers := strings.Split(line, ",")
		mid := 0
		ok := true
		visited := make(map[int]bool, 0)

		for i := 0; ok && i < len(numbers); i++ {
			num, _ := strconv.Atoi(numbers[i])

			visited[num] = true
			for _, shouldNotHaveSeen := range after[num] {
				if visited[shouldNotHaveSeen] {
					ok = false
					break
				}
			}

			// Middle number
			if i == len(numbers)/2 {
				mid = num
			}
		}

		if ok {
			acc += mid
		}
	}

	return acc
}
