package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/jacksehr/advent/utils"
)

func main() {
	half := utils.ParseArgs()

	scanner := utils.ScanFile("inputs/3.input")

	switch *half {
	case "first":
		solveFirstHalf(scanner)
	case "second":
		solveSecondHalf(scanner)
	}
}

func solveFirstHalf(scanner *bufio.Scanner) {
	var bitCounts = [][]int{}

	for scanner.Scan() {
		curr := scanner.Text()

		bits := strings.Split(curr, "")

		for len(bits) != len(bitCounts) {
			bitCounts = append(bitCounts, []int{0, 0})
		}

		for i, bitChar := range bits {
			bit, err := strconv.Atoi(bitChar)
			utils.Check(err)

			bitCounts[i][bit]++
		}
	}

	parseRates(bitCounts)
}

func solveSecondHalf(scanner *bufio.Scanner) {

}

func parseRates(input [][]int) {
	var gammaRateString, epsilonRateString string

	for _, count := range input {
		if count[0] > count[1] {
			gammaRateString += "0"
			epsilonRateString += "1"
		} else {
			gammaRateString += "1"
			epsilonRateString += "0"
		}
	}

	gammaRate, err := strconv.ParseInt(gammaRateString, 2, 0)
	utils.Check(err)

	epsilonRate, err := strconv.ParseInt(epsilonRateString, 2, 0)
	utils.Check(err)

	fmt.Println(gammaRate, epsilonRate, gammaRate*epsilonRate)
}
