package day03

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func PartASum(line string) int {
	acc := 0

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		acc += num1 * num2
	}

	return acc
}

func PartA(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	acc := 0

	for scanner.Scan() {
		line := scanner.Text()
		acc += PartASum(line)
	}

	return acc
}
