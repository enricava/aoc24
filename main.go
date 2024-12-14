package main

import (
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	// Change to current day
	day "github.com/enricava/aoc24/template"
)

func main() {
	if len(os.Args) != 4 {
		panic("Usage: <debug|release> <a|b> XX")
	}

	debug := os.Args[1] == "debug"
	part := os.Args[2]
	day := os.Args[3]

	initLogs(debug)

	t := time.Now()
	log.WithFields(log.Fields{
		"Solution": aoc(part, day),
		"Time":     time.Since(t),
	}).Info("Finished")
}

func aoc(part, day string) int {
	log.WithFields(log.Fields{
		" Part":       part,
		" Input file": day,
	}).Info("Running")

	file, err := openInputFile(day)
	if err != nil {
		log.WithFields(log.Fields{
			"input": day,
			"error": err,
		}).Fatal("Bad file input")
	}
	defer file.Close()

	runner := getRunner(part)
	if runner == nil {
		log.WithFields(log.Fields{
			"part": part,
		}).Fatal("Bad part input")
	}

	return runner(file)
}

func openInputFile(digits string) (*os.File, error) {
	input := "day" + digits + "/input.txt"
	return os.Open(input)
}

func getRunner(part string) func(file io.Reader) int {
	switch part {
	case "a":
		return day.PartA
	case "b":
		return day.PartB
	}
	return nil
}

func initLogs(debug bool) {
	log.SetOutput(os.Stdout)

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Debug("Running in debug mode")
}
