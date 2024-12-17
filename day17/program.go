package day17

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	adv uint8 = iota // A DIV 2^COMBO     -> A
	bxl              // B XOR LITERAL     -> B
	bst              // COMBO MOD 8       -> B
	jnz              // A != 0 jump to LITERAL, NO POINTER INCR
	bxc              // B XOR C           -> B
	out              // COMBO MOD 8       -> print
	bdv              // like adv          -> B
	cdv              // like adv          -> C
)

type Code struct {
	instructions []uint8
}

type Program struct {
	pc int

	RA int
	RB int
	RC int

	code Code
}

func logProgramState(p *Program) *log.Entry {
	return log.WithFields(log.Fields{
		"pc": p.pc,
		"RA": p.RA,
		"RB": p.RB,
		"RC": p.RC,
	})
}

func logCode(p *Program) *log.Entry {
	return log.WithFields(log.Fields{
		"instructions": p.code.instructions,
	})
}

func readRegister(scanner *bufio.Scanner, re *regexp.Regexp) int {
	scanner.Scan()
	line := scanner.Text()

	matches := re.FindStringSubmatch(line)
	register, _ := strconv.Atoi(matches[1])

	return register
}

func readCode(scanner *bufio.Scanner) Code {
	code := Code{[]uint8{}}

	scanner.Scan()
	line := scanner.Text()

	numbersString := strings.TrimPrefix(line, "Program: ")
	numbersSplit := strings.Split(numbersString, ",")

	for _, numStr := range numbersSplit {
		val, _ := strconv.Atoi(numStr)

		code.instructions = append(code.instructions, uint8(val))
	}

	return code
}

func readProgram(scanner *bufio.Scanner) *Program {
	var p Program

	registerRe := regexp.MustCompile(`Register .: (\d+)`)

	p.RA = readRegister(scanner, registerRe)
	p.RB = readRegister(scanner, registerRe)
	p.RC = readRegister(scanner, registerRe)

	// empty line
	scanner.Scan()

	p.code = readCode(scanner)

	return &p
}

func nextInstruction(p *Program) uint8 {
	return p.code.instructions[p.pc]
}

func nextLiteral(p *Program) int {
	return int(p.code.instructions[p.pc+1])
}

func nextCombo(p *Program) int {
	literal := nextLiteral(p)

	if literal <= 3 {
		return literal
	}

	switch literal {
	case 4:
		return p.RA
	case 5:
		return p.RB
	case 6:
		return p.RC
	}

	logProgramState(p).Warn("Program state")
	log.Warnf("Unexpected operand: %v", literal)

	return 0
}

func hasFinished(p *Program) bool {
	return !(p.pc < len(p.code.instructions)-1)
}

func process(p *Program) (isOutput bool, val int, finished bool) {
	instr := nextInstruction(p)
	combo := nextCombo(p)
	literal := nextLiteral(p)
	pcIncr := true

	switch instr {
	case adv:
		p.RA = p.RA / (1 << combo)
	case bxl:
		p.RB = p.RB ^ literal
	case bst:
		p.RB = combo & 7
	case jnz:
		if p.RA != 0 {
			pcIncr = false
			p.pc = int(literal)
		}
	case bxc:
		p.RB = p.RB ^ p.RC
	case out:
		isOutput = true
		val = combo % 8
	case bdv:
		p.RB = p.RA / (1 << combo)
	case cdv:
		p.RC = p.RA / (1 << combo)
	default:
		log.Fatal("Unimplemented instruction")
	}

	if pcIncr {
		p.pc += 2
	}

	finished = hasFinished(p)
	return
}
