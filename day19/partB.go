package day19

import (
	"io"

	log "github.com/sirupsen/logrus"
)

func PartB(file io.Reader) int {
	patterns, designs := readInput(file)

	for _, p := range patterns {
		if len(p) == 0 {
			panic("huih")
		}
	}

	sol := 0
	for _, d := range designs {
		valid := make(map[string]int)
		valid[""] = 1
		log.Debugf("Attempting: %s", d)
		validDesign(patterns, d, valid)
		sol += valid[d]
	}

	return sol
}
