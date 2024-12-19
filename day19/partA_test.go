package day19

import (
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestA(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	input := strings.NewReader(
		`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`,
	)

	result := PartA(input)
	expected := 6

	if expected != result {
		log.WithFields(log.Fields{
			"expected": expected,
			"result":   result,
		}).Warn("Failed")

		t.Fail()
	}
}
