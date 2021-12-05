package utils

import (
	"bufio"
	"flag"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ScanFile(filepath string) *bufio.Scanner {
	fileReader, err := os.Open(filepath)
	Check(err)

	return bufio.NewScanner(fileReader)
}

func ParseArgs() *string {
	solveHalf := flag.String("half", "first", "determine which part of the question to solve for (first or second half)")

	flag.Parse()

	return solveHalf
}
