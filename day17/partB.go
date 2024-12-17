package day17

import (
	"bufio"
	"io"
	"reflect"

	log "github.com/sirupsen/logrus"
)

func PartB(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	p := readProgram(scanner)

	sol := 0

	logProgramState(p).Debug("Read program")
	logCode(p).Debug("Read code")

	for i := 1000000000; i < 10000000000; i++ {
		output := make([]int, 0)

		var rep Program
		rep = *p
		p.RA = i
		for {
			isOutput, val, finished := process(&rep)

			if isOutput {
				output = append(output, val)
			}

			if finished {
				break
			}
		}

		if reflect.DeepEqual(p.code.instructions, output) {
			sol = i
			break
		}
		log.Debug(p.code)
		log.Debug(output)
	}

	return sol
}
