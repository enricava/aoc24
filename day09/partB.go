package day09

import (
	"bufio"
	"io"
	"strconv"
)

func compactChecksumNoFragmentation(fileLengths, paddingLengths []int) int {
	checksum := 0

	offset := fileLengths[0]
	fileLengths[0] = 0 // first file is consumed
	paddingStarts := make([]int, len(paddingLengths))
	fileStarts := make([]int, len(fileLengths))

	// initialize padding and file starts
	for i := 1; i <= len(paddingLengths); i++ {
		paddingStarts[i-1] = offset
		offset += paddingLengths[i-1]
		fileStarts[i] = offset
		offset += fileLengths[i]
	}

	for file := len(fileLengths) - 1; file > 0; file-- {

		// find first padding that can fit the file
		padding := 0
		for padding < file {
			if paddingLengths[padding] >= fileLengths[file] {
				break
			}
			padding++
		}
		if paddingLengths[padding] < fileLengths[file] {
			continue
		}

		// fit the file
		consumed := fileLengths[file]
		start := paddingStarts[padding]
		end := paddingStarts[padding] + consumed - 1

		checksum += sumFromTo(start, end) * file

		paddingLengths[padding] -= consumed
		fileLengths[file] -= consumed
		paddingStarts[padding] += consumed
	}

	// not moved files
	for file, length := range fileLengths {
		if length < 1 {
			continue
		}

		consumed := fileLengths[file]
		start := fileStarts[file]
		end := fileStarts[file] + consumed - 1

		checksum += sumFromTo(start, end) * file
	}

	return checksum
}

func PartB(file io.Reader) int {
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

	return compactChecksumNoFragmentation(fileLengths, paddings)
}
