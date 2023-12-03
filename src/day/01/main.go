package day01

import (
	"advent-of-code-2023/utils"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

func RunA(input []string) (int, error) {
	if input == nil {
		input = utils.ReadLines("./day/01/input.txt")
	}

	digitRegexp := regexp.MustCompile("\\d")

	values := funk.Map(input, func(line string) int {
		digits := digitRegexp.FindAllString(line, -1)
		lineValue := digits[0] + funk.Last(digits).(string)
		lineNumber, error := strconv.Atoi(lineValue)
		if error != nil {
			log.Fatal(error)
		}
		return lineNumber
	}).([]int)

	solution := funk.SumInt(values)
	return solution, nil
}

func RunB(input []string) (int, error) {
	if input == nil {
		input = utils.ReadLines("./day/01/input.txt")
	}

	values := funk.Map(input, func(line string) int {

		firstDigit, lastDigit := findExternalDigits(line)
		lineNumber, error := strconv.Atoi(firstDigit + lastDigit)
		if error != nil {
			log.Fatal(error)
		}

		return lineNumber
	}).([]int)

	solution := funk.SumInt(values)
	return solution, nil
}

var cardinalValues = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"ten":   "10",
}

var digitRegexp = regexp.MustCompile("\\d")

type cardinalPosition struct {
	index int
	value string
}

func findExternalDigits(value string) (string, string) {
	allDigitIndexes := digitRegexp.FindAllStringIndex(value, -1)

	var firstCardinal cardinalPosition = cardinalPosition{}
	var lastCardinal cardinalPosition = cardinalPosition{}

	for cardinal, cardinalValue := range cardinalValues {

		cardinalFirstIndex := strings.Index(value, cardinal)
		cardinalLastIndex := strings.LastIndex(value, cardinal)

		if cardinalFirstIndex != -1 &&
			(firstCardinal == cardinalPosition{} || firstCardinal.index > cardinalFirstIndex) {
			firstCardinal = cardinalPosition{index: cardinalFirstIndex, value: cardinalValue}
		}
		if cardinalLastIndex != -1 &&
			(lastCardinal == cardinalPosition{} || lastCardinal.index < cardinalLastIndex) {
			lastCardinal = cardinalPosition{index: cardinalLastIndex, value: cardinalValue}
		}
	}

	if (len(allDigitIndexes) <= 0 && firstCardinal == cardinalPosition{}) {
		log.Fatal("No digit or cardinal values found: \"" + value + "\"")
	}

	var firstDigit string = ""
	var lastDigit string = ""

	digitsLength := len(allDigitIndexes)
	if digitsLength > 0 {
		if (firstCardinal == cardinalPosition{} || allDigitIndexes[0][0] < firstCardinal.index) {
			firstDigit = string(value[allDigitIndexes[0][0]])
		} else {
			firstDigit = firstCardinal.value
		}
		if (lastCardinal == cardinalPosition{} || allDigitIndexes[digitsLength-1][0] > lastCardinal.index) {
			lastDigit = string(value[allDigitIndexes[digitsLength-1][0]])
		} else {
			lastDigit = lastCardinal.value
		}
	} else {
		firstDigit = firstCardinal.value
		lastDigit = lastCardinal.value
	}

	return firstDigit, lastDigit
}
