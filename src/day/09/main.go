package day09

import (
	"advent-of-code-2023/utils"
	"log"

	"github.com/thoas/go-funk"
)

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadLines(inputFile)
	} else {
		return utils.ReadLines("./day/09/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	input := getInput(inputFile)
	readings := parseReadings(input)

	total := 0
	for _, reading := range readings {
		nextValue := findNextSequenceValue(reading)
		total += nextValue
	}

	return total, nil
}

func RunB(inputFile string) (int, error) {
	input := getInput(inputFile)
	readings := parseReadings(input)

	total := 0
	for _, reading := range readings {
		previousValue := findPreviousSequenceValue(reading)
		total += previousValue
	}

	return total, nil
}
func parseReadings(input []string) [][]int {
	readings := [][]int{}
	for _, line := range input {
		values, _ := utils.StringToInts(line)
		readings = append(readings, values)

	}
	return readings
}

func findNextSequenceValue(readings []int) int {
	differences := findDifferenceSequences(readings)
	var sequence []int

	for i := len(differences) - 1; i > 0; i-- {
		sequence = differences[i]
		nextSequence := &differences[i-1]

		nextValue := funk.Last(sequence).(int) + funk.Last(nextSequence).(int)
		*nextSequence = append(*nextSequence, nextValue)
	}

	return funk.Last(differences[0]).(int)
}

func findPreviousSequenceValue(readings []int) int {
	differences := findDifferenceSequences(readings)
	var sequence []int

	for i := len(differences) - 1; i > 0; i-- {
		sequence = differences[i]
		nextSequence := &differences[i-1]

		nextValue := differences[i-1][0] - sequence[0]
		*nextSequence = append([]int{nextValue}, *nextSequence...)
	}

	return differences[0][0]
}

func findDifferenceSequences(readings []int) [][]int {
	differences := [][]int{readings}
	sequence := readings

	for true {
		diff := findDifferenceSequence(sequence)
		differences = append(differences, diff)

		nonZeros := funk.FilterInt(diff, func(value int) bool {
			return value != 0
		})

		if len(nonZeros) == 0 {
			return differences
		}

		if len(diff) < 2 {
			log.Fatal("Failed to extrapolate next value")
			return [][]int{}
		}

		sequence = diff
	}

	return [][]int{}
}

func findDifferenceSequence(readings []int) []int {
	leni := len(readings) - 1
	diff := make([]int, leni)

	for i := 0; i < leni; i++ {
		reading := readings[i]
		nextReading := readings[i+1]
		diff[i] = nextReading - reading
	}

	return diff
}
