package day02

import (
	"bufio"
	"io"
)

func skipOne(slice []int, skip_idx int) (newSlice []int) {
	newSlice = make([]int, 0)
	newSlice = append(newSlice, slice[:skip_idx]...)
	newSlice = append(newSlice, slice[skip_idx+1:]...)
	return
}

func inRange(slice []int, idx int) bool {
	return idx >= 0 && idx < len(slice)
}

func IsMostlySafe(nums []int, badIndex int) bool {
	ok := false
	skipIndexes := []int{badIndex - 2, badIndex - 1, badIndex}

	for i := 0; !ok && i < 3; i++ {
		skipIndex := skipIndexes[i]

		if inRange(nums, skipIndex) {
			ok, _ = IsSafe(skipOne(nums, skipIndex))
		}
	}

	return ok
}

func PartB(file io.Reader) any {
	scanner := bufio.NewScanner(file)

	acc := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := parseInts(line)

		if ok, badIndex := IsSafe(nums); ok || IsMostlySafe(nums, badIndex) {
			acc++
		}
	}

	return acc
}
