package day01

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

func readListCount(file io.Reader) (list[]int, count map[int]int) {
	reader := bufio.NewReader(file)

	re := regexp.MustCompile(`^(\d+)\s+(\d+)`)

	list = make([]int, 0)
	count = make(map[int]int, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		match := re.FindStringSubmatch(line)
		if match != nil {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			list = append(list, num1)
      _, ok := count[num2]
      if !ok {
        count[num2] = 0
      }

      count[num2]++
		}
	}

  return
}

func PartB(file io.Reader) any{
  list, count := readListCount(file)

  acc := 0
  for _, val := range list {
    acc += val * count[val]
  }

  return acc
}
