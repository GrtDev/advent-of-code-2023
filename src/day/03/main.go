package day03

import (
	"advent-of-code-2023/utils"
	"regexp"
	"strconv"
	"strings"
)

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadParagraphs(inputFile)
	} else {
		return utils.ReadParagraphs("./day/05/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	var input []string = getInput(inputFile)

	matrix2d := parseInput(input)
	totalValue := 0
	for y := range matrix2d {
		totalValue += getLineValue(y, matrix2d)
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

	matrix2d := parseInput(input)
	totalValue := 0
	gearRegexp := regexp.MustCompile("\\*")

	for y := range matrix2d {
		line := strings.Join(matrix2d[y], "")
		gearIndexes := gearRegexp.FindAllStringIndex(line, -1)

		for _, gearIndex := range gearIndexes {
			gearNumbers := getGearNumbers(point{gearIndex[0], y}, matrix2d)
			if len(gearNumbers) == 2 {
				totalValue += gearNumbers[0] * gearNumbers[1]
			}
		}
	}

	return totalValue, nil
}

type point struct {
	x int
	y int
}

func parseInput(input []string) [][]string {
	matrix2d := [][]string{}
	for _, line := range input {
		matrix2d = append(matrix2d, strings.Split(line, ""))
	}
	return matrix2d
}

func getLineValue(y int, matrix [][]string) int {
	digitsRegexp := regexp.MustCompile("\\d+")
	line := strings.Join(matrix[y], "")
	numberIndexes := digitsRegexp.FindAllStringIndex(line, -1)

	totalValue := 0
	for _, indexes := range numberIndexes {
		start := indexes[0]
		end := indexes[1]

		if (indexRangeTouchesSymbol(point{start, y}, end-start, matrix)) {
			numberValue, _ := strconv.Atoi(line[start:end])
			totalValue += numberValue
		}
	}

	return totalValue
}

func indexRangeTouchesSymbol(startPos point, xSteps int, matrix [][]string) bool {
	for x := startPos.x; x < startPos.x+xSteps; x++ {
		if (positionTouchesSymbol(point{x, startPos.y}, matrix)) {
			return true
		}
	}
	return false
}

func positionTouchesSymbol(pos point, matrix [][]string) bool {
	gridSize := 3
	gridValues := [9]string{}
	startPosition := point{x: pos.x - 1, y: pos.y - 1}

	for yMove := 0; yMove < gridSize; yMove++ {
		for xMove := 0; xMove < gridSize; xMove++ {
			testPos := point{startPosition.x + xMove, startPosition.y + yMove}
			gridValues[yMove*gridSize+xMove] = safeGet(testPos, matrix)
		}
	}

	var symbolRegexp = regexp.MustCompile("[^.\\d]")
	for _, v := range gridValues {
		if symbolRegexp.MatchString(v) {
			return true
		}
	}

	return false
}

func safeGet(pos point, matrix [][]string) string {
	if pos.y < 0 || pos.y >= len(matrix) {
		return ""
	}
	line := matrix[pos.y]
	if pos.x < 0 || pos.x >= len(line) {
		return ""
	}
	return line[pos.x]
}

func getGearNumbers(gearPos point, matrix [][]string) []int {
	digitsRegexp := regexp.MustCompile("\\d+")
	gridSize := 3
	startPosition := point{x: gearPos.x - 1, y: gearPos.y - 1}

	gearNumbers := []int{}

	for y := startPosition.y; y < startPosition.y+gridSize; y++ {
		if y < 0 || y >= len(matrix) {
			continue
		}
		line := strings.Join(matrix[y], "")
		numberIndexes := digitsRegexp.FindAllStringIndex(line, -1)

		for _, indexes := range numberIndexes {
			start := indexes[0]
			end := indexes[1]
			if start <= gearPos.x+1 && end > gearPos.x-1 {
				numberValue, _ := strconv.Atoi(line[start:end])
				gearNumbers = append(gearNumbers, numberValue)
			}
		}
	}

	return gearNumbers
}
