package day03

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func PartB(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	var builder strings.Builder

	acc := 0

  /// XXX: This is a cool hack to avoid leaking state
	for scanner.Scan() {
		line := scanner.Text()
		builder.WriteString(line)
	}

	text := builder.String()

	do := `do()`
	dont := `don't()`

	for len(text) > 0 {
		// advance up to dont
		left_from_dont, right_from_dont, _ := strings.Cut(text, dont)

		// within dont - do
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
		matches := re.FindAllStringSubmatch(left_from_dont, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			acc += num1 * num2
		}

		//advance up to do
		_, right_from_do, _ := strings.Cut(right_from_dont, do)
		text = right_from_do
	}

	return acc
}
