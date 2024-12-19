package day19

import (
	"bufio"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"
)

var fmax int

func equal(p, s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(p) > len(s) {
		return false
	}

	for i := 0; i < len(p); i++ {
		if p[i] != s[i] {
			return false
		}
	}
	return true
}

func validDesign(ps []string, d string, valid map[string]int) int {
	if len(d) == 0 {
		return 1
	}

	if val, exists := valid[d]; exists {
		return val
	}

	for _, p := range ps {
		if len(p) == 0 {
			continue
		}

		l := len(p)
		if l > len(d) {
			continue
		}

		s := d[:l]
		r := d[l:]

		if equal(p, s) {
			valid[d] += validDesign(ps, r, valid)
		}
	}

	return valid[d]
}

func readInput(file io.Reader) (patterns, designs []string) {
	scanner := bufio.NewScanner(file)

	// patterns
	scanner.Scan()

	line := scanner.Text()
	patterns = strings.Split(line, ", ")

	// dismiss empty line
	scanner.Scan()
	scanner.Text()

	// designs
	for scanner.Scan() {
		line = scanner.Text()
		designs = append(designs, line)
	}

	// log.WithFields(log.Fields{
	// 	"npatterns":      len(patterns),
	// 	"samplePatterns": patterns[:3],
	// 	"ndesigns":       len(designs),
	// 	"sampleDesigns":  designs[:3],
	// }).Debug("Got input")

	return
}

func PartA(file io.Reader) int {
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
		if valid[d] > 0 {
			log.Debugf("%s is valid", d)
			sol++
		}
	}

	return sol
}
