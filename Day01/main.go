package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var DigitsToLettters = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var (
	exp  = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	exp2 = regexp.MustCompile(".*(" + exp.String() + ")")
)

func main() {
	rows := parseInput("input.txt")

	var sum int
	var sum2 int

	for _, row := range rows {
		// Part 1
		firstNum := strings.IndexFunc(row, unicode.IsDigit)
		lastNum := strings.LastIndexFunc(row, unicode.IsDigit)
		v, err := strconv.Atoi(string(row[firstNum]) + string(row[lastNum]))
		if err != nil {
			fmt.Print(err)
		}
		sum += v

		// Part 2
		first2 := DigitsToLettters[exp.FindString(row)]
		last2 := DigitsToLettters[exp2.FindStringSubmatch(row)[1]]
		v2 := 10*first2 + last2
		sum2 += v2
	}
	fmt.Println(sum)
	fmt.Println(sum2)
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
