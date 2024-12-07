package day07

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

func numDigits(v int) int {
	n := 0
	for c := v; c > 0; c = c / 10 {
		n++
	}
	return max(n, 1)
}

func concat(a, b int) int {
	n := numDigits(b)
	return a*int(math.Pow10(n)) + b
}

func equatableB(res, i, cur int, nums []int) bool {
	if cur > res {
		return false
	}

	if i == len(nums) {
		return cur == res
	}

	equatableMult := equatableB(res, i+1, cur*nums[i], nums)
	equatableSum := equatableB(res, i+1, cur+nums[i], nums)
	equatableConcat := equatableB(res, i+1, concat(cur, nums[i]), nums)

	return equatableMult || equatableSum || equatableConcat
}

func PartB(file io.Reader) int {
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
		if equatableB(res, 1, cur, nums) {
			acc += res
		}
	}

	return acc
}
