package day02

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func sign(num int) int {
	if num > 0 {
		return 1
	} else if num < 0 {
		return -1
	}
	return 0
}

func parseInts(line string) []int {
	nums_str := strings.Fields(line)
	nums := make([]int, len(nums_str))

	for i, v := range nums_str {
		num, _ := strconv.Atoi(v)
		nums[i] = num
	}

	return nums
}

// Whether it is safe or the bad level
func IsSafe(nums []int) (bool, int) {
	ok := true
	last := 0
	sign := sign(nums[1] - nums[0])

	for i := 1; ok && i < len(nums); i++ {
		diff := sign * (nums[i] - nums[i-1])

		ok = diff >= 1 && diff <= 3
		last = i
	}

	return ok, last
}

func PartA(file io.Reader) any {
	scanner := bufio.NewScanner(file)

	acc := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := parseInts(line)

		if ok, _ := IsSafe(nums); ok {
			acc++
		}
	}

	return acc
}
