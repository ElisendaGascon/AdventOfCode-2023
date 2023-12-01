package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	rows := parseInput("input.txt")

	var sum int

	for _, row := range rows {
		firstNum := strings.IndexFunc(row, unicode.IsDigit)
		lastNum := strings.LastIndexFunc(row, unicode.IsDigit)
		v, err := strconv.Atoi(string(row[firstNum]) + string(row[lastNum]))
		if err != nil {
			fmt.Print(err)
		}
		sum += v
	}
	fmt.Print(sum)
}

func parseInput(inputFile string) []string {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Print(err)
	}
	inputString := string(input)

	rows := strings.Split(inputString, "\n")
	return rows
}
