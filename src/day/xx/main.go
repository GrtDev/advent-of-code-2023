package dayxx

import (
	"advent-of-code-2023/utils"
)

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadLines(inputFile)
	} else {
		return utils.ReadLines("./day/xx/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	input := getInput(inputFile)
	parsed := parseInput(input)

	total := 0
	for _, value := range parsed {
		total += value[0]
	}

	return total, nil
}

func RunB(inputFile string) (int, error) {
	return -1, nil
}

func parseInput(input []string) [][]int {
	parsedInput := [][]int{}
	for _, line := range input {
		values, _ := utils.StringToInts(line)
		parsedInput = append(parsedInput, values)
	}
	return parsedInput
}
