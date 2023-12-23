package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := parseInput("input.txt")

	maxColours := findMaxColours(lines)
	sum := 0
	sum2 := 0

	for idx, maxColour := range maxColours {
		// Part 1
		max := []int{14, 13, 12} // blue, green, red
		if maxColour[0] <= max[0] && maxColour[1] <= max[1] && maxColour[2] <= max[2] {
			sum += idx + 1
		}
		// Part 2
		val := maxColour[0] * maxColour[1] * maxColour[2]
		sum2 += val
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func findMaxColours(lines []string) [][]int {
	result := [][]int{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid input", parts)
		}

		sets := strings.Split(parts[1], ";")
		maxSet := []int{0, 0, 0}
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				draw := strings.Split(strings.TrimSpace(cube), " ")
				num, _ := strconv.Atoi(draw[0])
				colour := parseColour(draw[1])
				maxSet[colour] = max(maxSet[colour], num)
			}
		}
		result = append(result, maxSet)
	}
	return result
}

func parseInput(inputFile string) []string {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Print(err)
	}
	inputString := string(input)

	lines := strings.Split(inputString, "\n")
	return lines
}

func parseColour(colour string) int {
	switch colour {
	case "blue":
		return 0
	case "green":
		return 1
	case "red":
		return 2
	default:
		return -1
	}
}
