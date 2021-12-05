package main

import (
	"bufio"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/jacksehr/advent/utils"
)

const (
	Forward = "forward"
	Down    = "down"
	Up      = "up"
)

func parseArgs() *string {
	solveHalf := flag.String("half", "first", "determine which part of the question to solve for (first or second half)")

	flag.Parse()

	return solveHalf
}

func main() {
	half := parseArgs()

	scanner := utils.ScanFile("inputs/2.input")

	switch *half {
	case "first":
		solveFirstHalf(scanner)
	case "second":
		solveSecondHalf(scanner)
	}
}

func solveFirstHalf(scanner *bufio.Scanner) {
	var position = []int{0, 0} // hPos, vPos

	for scanner.Scan() {
		curr := scanner.Text()

		instruction := strings.Split(curr, " ")
		direction := instruction[0]
		distance, err := strconv.Atoi(instruction[1])

		utils.Check(err)

		switch direction {
		case Forward:
			position[0] += distance
		case Up:
			position[1] -= distance
		case Down:
			position[1] += distance
		}
	}

	fmt.Println("Final position product:", position[0]*position[1])
}

func solveSecondHalf(scanner *bufio.Scanner) {
	var position = []int{0, 0} // hPos, vPos
	var aim int

	for scanner.Scan() {
		curr := scanner.Text()

		instruction := strings.Split(curr, " ")
		direction := instruction[0]
		distance, err := strconv.Atoi(instruction[1])

		utils.Check(err)

		switch direction {
		case Down:
			aim += distance
		case Up:
			aim -= distance
		case Forward:
			position[0] += distance
			position[1] += aim * distance
		}
	}

	fmt.Println("Final position product:", position[0]*position[1])
}
