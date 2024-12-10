package day09

import (
	"bufio"
	"io"
	"strconv"
)

func sumFromTo(m, n int) int {
	return ((m + n) * (n - m + 1)) / 2
}

// This is quite fast! (300 us)
func compactChecksum(fileLengths, paddings []int) int {
	checksum := 0

	offset := fileLengths[0] // skip idx = 0 in checksum
	fileLengths[0] = 0       // first file is consumed
	paddingIdx := 0

	// fmt.Println("files:", fileLengths, "| paddings:", paddings, "| offset:",
	// 	offset, "| paddingIdx:", paddingIdx)

	for fileIdx := len(fileLengths) - 1; fileIdx > 0; fileIdx-- {

		// how many blocks from the padding where consumed
		consumed := min(paddings[paddingIdx], fileLengths[fileIdx])
		paddings[paddingIdx] -= consumed
		fileLengths[fileIdx] -= consumed

		checksum += sumFromTo(offset, offset+consumed-1) * fileIdx
		offset += consumed

		// fmt.Println("consumed:", consumed, "| fileLengths:", fileLengths, "| paddings:",
		// 	paddings, "| offset:", offset, "| paddingIdx:", paddingIdx, "| checksum:", checksum)

		if paddingIdx < len(paddings) && paddings[paddingIdx] == 0 {
			paddingIdx++

			// padding consumed, add file from the left
			consumed = fileLengths[paddingIdx]
			fileLengths[paddingIdx] = 0

			checksum += sumFromTo(offset, offset+consumed-1) * paddingIdx
			offset += consumed
		}

		// keep going for this file
		if fileLengths[fileIdx] > 0 {
			fileIdx++
		}
	}

	return checksum
}

func PartA(file io.Reader) int {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	fileLengths := []int{}
	paddings := []int{}

	for i, v := range line {
		num, _ := strconv.Atoi(string(v))

		if i%2 == 0 {
			fileLengths = append(fileLengths, num)
		} else {
			paddings = append(paddings, num)
		}
	}

	return compactChecksum(fileLengths, paddings)
}
