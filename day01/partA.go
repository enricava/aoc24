package day01

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strconv"
)

func readPairs(file io.Reader) (list1, list2 []int) {
	reader := bufio.NewReader(file)

	re := regexp.MustCompile(`^(\d+)\s+(\d+)`)

	list1 = make([]int, 0)
	list2 = make([]int, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		match := re.FindStringSubmatch(line)
		if match != nil {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

  return
}

func abs(x int) int {
  if (x < 0) {
    return -x
  }
  return x
}

func PartA(file io.Reader) any {
  list1, list2 := readPairs(file)

  sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j]})
  sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j]})

  acc := 0
  for i := 0; i < len(list1); i++ {
    acc += abs(list1[i] - list2[i])
  }

  return acc
}
