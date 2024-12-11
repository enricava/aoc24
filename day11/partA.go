package day11

import (
	"bufio"
	"fmt"
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

func split(num, at int) (a, b int) {
	div := int(math.Pow10(at))

	a = num / div
	b = num % div

	return
}

func printStones(stones map[int]int) {
	for k, v := range stones {
		fmt.Println(k, v)
	}
}

func countStones(stones map[int]int) int {
	acc := 0
	for _, count := range stones {
		acc += count
	}
	return acc
}

func generateStones(stones map[int]int) map[int]int {
	newStones := make(map[int]int)

	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
		} else if digits := numDigits(stone); digits%2 == 0 {
			stoneA, stoneB := split(stone, digits/2)
			newStones[stoneA] += count
			newStones[stoneB] += count
		} else {
			newStones[stone*2024] += count
		}
	}

	return newStones
}

func stonesAfterNBlinks(stones map[int]int, n int) int {
	for i := 0; i < n; i++ {
		stones = generateStones(stones)
	}

	return countStones(stones)
}

func inputToStones(file io.Reader) map[int]int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	stones_str := strings.Split(line, " ")
	stones := make(map[int]int)
	for _, v := range stones_str {
		stone, _ := strconv.Atoi(v)
		stones[stone]++
	}

	return stones
}

func PartA(file io.Reader) int {
	stones := inputToStones(file)
	return stonesAfterNBlinks(stones, 25)
}
