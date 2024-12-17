package day17

import (
	"bufio"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestA(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	input := strings.NewReader(
		`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
	)

	PartA(input)
	println("expected 4,6,3,5,6,3,5,2,1,0")
}

func TestA1(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	input := strings.NewReader(
		`Register A: 0
Register B: 0
Register C: 9

Program: 2,6`,
	)

	scanner := bufio.NewScanner(input)
	p := readProgram(scanner)
	process(p)

	if p.RB != 1 {
		t.Fail()
	}
}

func TestA2(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	input := strings.NewReader(
		`Register A: 10
Register B: 0
Register C: 0

Program: 5,0,5,1,5,4`,
	)

	PartA(input)
	println("expected 0,1,2")

}

func TestA3(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	input := strings.NewReader(
		`Register A: 2024
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
	)

	PartA(input)
	println("expected 4,2,5,6,7,7,7,7,3,1,0")
}

func TestA4(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	input := strings.NewReader(
		`Register A: 0
Register B: 29
Register C: 0

Program: 1,7`,
	)

	scanner := bufio.NewScanner(input)
	p := readProgram(scanner)
	process(p)

	if p.RB != 26 {
		t.Fail()
	}
}

func TestA5(t *testing.T) {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	input := strings.NewReader(
		`Register A: 0
Register B: 2024
Register C: 43690

Program: 4,0`,
	)

	scanner := bufio.NewScanner(input)
	p := readProgram(scanner)
	process(p)

	if p.RB != 44354 {
		t.Fail()
	}
}
