package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jacksehr/advent/utils"
)

func main() {
	half := utils.ParseArgs()

	scanner := utils.ScanFile("inputs/1.input")

	switch *half {
	case "first":
		solveFirstHalf(scanner)
	case "second":
		solveSecondHalf(scanner)
	}
}

func solveFirstHalf(scanner *bufio.Scanner) {
	var prev, count int
	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		utils.Check(err)

		if prev != 0 && curr > prev {
			count++
		}

		prev = curr
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read:", err)
	}

	fmt.Println("Number of increases:", count)
}

func solveSecondHalf(scanner *bufio.Scanner) {
	var index, count, sumA, sumB int
	var window = []int{0, 0, 0, 0}

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		utils.Check(err)

		if index < 4 {
			window[index] = curr

			index++

			if index != 4 {
				continue
			}
		} else {
			window = append(window[1:4], curr)
		}

		sumA, sumB = 0, 0

		// sumA
		for _, i := range window[0:3] {
			sumA += i
		}

		// sumB
		for _, i := range window[1:4] {
			sumB += i
		}

		if sumB > sumA {
			count++
		}

		index++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read:", err)
	}

	fmt.Println("Number of increases:", count)
}
