package day07

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func equatableA(res, i, cur int, nums []int) bool {
	if cur > res {
		return false
	}

	if i == len(nums) {
		return cur == res
	}

	equatableMult := equatableA(res, i+1, cur*nums[i], nums)
	equatableSum := equatableA(res, i+1, cur+nums[i], nums)

	return equatableMult || equatableSum
}

func PartA(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	acc := 0

	for scanner.Scan() {
		line := scanner.Text()
		left, right, _ := strings.Cut(line, ":")

		// Left side number
		res, _ := strconv.Atoi(left)

		// Right side numbers
		right = strings.TrimSpace(right)
		numStrings := strings.Split(right, " ")
		nums := make([]int, len(numStrings))
		for i, numstr := range numStrings {
			nums[i], _ = strconv.Atoi(numstr)
		}

		// Calc
		cur := nums[0]
		if equatableA(res, 1, cur, nums) {
			acc += res
		}
	}

	return acc
}
