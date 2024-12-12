package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	// Change to current day
	day "github.com/enricava/aoc24/day12"
)

func main() {
	if len(os.Args) != 3 {
		panic("Usage: <a|b> XX")
	}

	part := os.Args[1]
	day := os.Args[2]

	fmt.Println(" Part:", part)
	fmt.Println(" Input:", day)
	fmt.Println("---")

	aoc(part, day)
}

func aoc(part, digits string) {
	input := "day" + digits + "/input.txt"
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var runner func(file io.Reader) int
	switch part {
	case "a":
		runner = day.PartA
	case "b":
		runner = day.PartB
	}

	defer func(t time.Time) {
		fmt.Println("Execution took:", time.Since(t))
	}(time.Now())

	result := runner(file)
	fmt.Println(result)
}
