package day17

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	log "github.com/sirupsen/logrus"
)

func PartA(file io.Reader) int {
	scanner := bufio.NewScanner(file)

	p := readProgram(scanner)

	logProgramState(p).Debug("Read program")
	logCode(p).Debug("Read code")

	var sb strings.Builder


	for {
		isOutput, val, finished := process(p)

		if isOutput {
			sb.WriteString(fmt.Sprintf("%v,", val))
		}

		if finished {
			break
		}
	}

  logProgramState(p).Debug("Finished")

	log.Info(sb.String())

	sol := 0
	return sol
}
