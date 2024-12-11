package day11

import (
	"io"
)


func PartB(file io.Reader) int {
	stones := inputToStones(file)
	return stonesAfterNBlinks(stones, 75)
}
