package day16

import (
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestA(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	input := strings.NewReader(
		`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`,
	)

	result := PartA(input)
	expected := 7036

	if expected != result {
		log.WithFields(log.Fields{
			"expected": expected,
			"result":   result,
		}).Warn("Failed")

		t.Fail()
	}
}

func TestA2(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	input := strings.NewReader(
		`#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`,
	)

	result := PartA(input)
	expected := 11048

	if expected != result {
		log.WithFields(log.Fields{
			"expected": expected,
			"result":   result,
		}).Warn("Failed")

		t.Fail()
	}
}
