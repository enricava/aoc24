package day03

import (
	"bufio"
	"io"
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
		left_from_dont, right_from_dont, _ := strings.Cut(text, dont)

		// within dont - do
		acc += PartASum(left_from_dont)

		_, right_from_do, _ := strings.Cut(right_from_dont, do)
		text = right_from_do
	}

	return acc
}
