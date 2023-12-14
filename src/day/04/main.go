package day04

import (
	"advent-of-code-2023/utils"
	"regexp"

	"github.com/thoas/go-funk"
)

type card struct {
	winning    []int
	numbers    []int
	numbersWon []int
	numCopies  int
}

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadParagraphs(inputFile)
	} else {
		return utils.ReadParagraphs("./day/05/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	var input []string = getInput(inputFile)

	cards := parseInput(input)

	totalValue := 0
	for _, card := range cards {
		lenNumbersWon := len(card.numbersWon)

		if lenNumbersWon > 0 {
			totalValue += utils.PowInt(2, lenNumbersWon-1)
		}
	}

	return totalValue, nil
}

func RunB(inputFile string) (int, error) {
	var input []string
	if inputFile != "" {
		input = utils.ReadLines(inputFile)
	} else {
		input = utils.ReadLines("./input.txt")
	}

	cards := parseInput(input)
	lenCards := len(cards)

	totalCards := 0
	for i, card := range cards {
		lenNumbersWon := len(card.numbersWon)
		totalCards += card.numCopies

		for j := i + 1; j < funk.MinInt([]int{i + lenNumbersWon + 1, lenCards}); j++ {
			cards[j].numCopies += card.numCopies
		}
	}

	return totalCards, nil
}

func parseInput(input []string) []card {
	cardRegexp := regexp.MustCompile(`^Card\s+\d+:\s+([\d\s]+)\|([\d\s]+)$`)

	cards := []card{}
	for _, line := range input {
		matches := cardRegexp.FindStringSubmatch(line)
		if len(matches) > 0 {
			winning, errWinning := utils.StringToInts(matches[1], " ")
			numbers, errNumbers := utils.StringToInts(matches[2], " ")

			numbersWon := funk.FilterInt(numbers, func(number int) bool {
				return funk.ContainsInt(winning, number)
			})

			utils.FatalOnError(errWinning)
			utils.FatalOnError(errNumbers)

			cards = append(cards, card{
				winning:    winning,
				numbers:    numbers,
				numbersWon: numbersWon,
				numCopies:  1,
			})
		}
	}
	return cards
}
