package day19

import (
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestB(t *testing.T) {
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

	result := PartB(input)
	expected := 16

	if expected != result {
		log.WithFields(log.Fields{
			"expected": expected,
			"result":   result,
		}).Warn("Failed")

		t.Fail()
	}
}
