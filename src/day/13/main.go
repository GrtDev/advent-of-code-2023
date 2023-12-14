package day13

import (
	"advent-of-code-2023/utils"
	"advent-of-code-2023/utils/matrix"
	"fmt"
	"log"
)

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadParagraphs(inputFile)
	} else {
		return utils.ReadParagraphs("./day/13/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	input := getInput(inputFile)
	notes := parseInput(input)

	total := 0
	for _, note := range notes {
		x, _, y, _ := findMirrorIndex(note, -1, -1)
		if x == -1 && y == -1 {
			log.Fatal("failed to find mirror index")
		}

		if y >= 0 {
			total += (y + 1) * 100
		}
		if x >= 0 {
			total += x + 1
		}
	}

	return total, nil
}

func RunB(inputFile string) (int, error) {
	input := getInput(inputFile)
	notes := parseInput(input)

	total := 0
	for _, note := range notes {
		mirrorX, mirrorRangeX, mirrorY, mirrorRangeY := findMirrorIndex(note, -1, -1)
		x, _, y, _ := findMirrorIndexWithSmudge(note, mirrorX, mirrorRangeX, mirrorY, mirrorRangeY)

		if x == -1 && y == -1 {
			fmt.Printf(" matrix: \n%v\n", matrix.Print(note))
			log.Fatal("failed to find mirror index")
		}

		if y >= 0 && y != mirrorY {
			total += (y + 1) * 100
		}
		if x >= 0 && x != mirrorX {
			total += x + 1
		}
	}

	return total, nil
}

func findMirrorIndexWithSmudge(note [][]string, mirrorX int, mirrorRangeX []int, mirrorY int, mirrorRangeY []int) (int, []int, int, []int) {
	newRangeX, newRangeY := []int{-1, -1}, []int{-1, -1}
	newX, newY := -1, -1

	if mirrorX != -1 {
		if mirrorRangeX[0] > 0 {
			for x := 0; x < mirrorRangeX[0]; x++ {
				newX, newRangeX, newY, newRangeY = removeSmudgeColumn(note, x, mirrorX, mirrorY)
				if newX != -1 && newX != mirrorX || newY != -1 && newY != mirrorY {
					return newX, newRangeX, newY, newRangeY
				}
			}
		} else {
			for x := mirrorRangeX[1]; x < matrix.LenRow(note); x++ {
				newX, newRangeX, newY, newRangeY = removeSmudgeColumn(note, x, mirrorX, mirrorY)
				if newX != -1 && newX != mirrorX || newY != -1 && newY != mirrorY {
					return newX, newRangeX, newY, newRangeY
				}
			}
		}
	}

	if mirrorY != -1 {
		if mirrorRangeY[0] > 0 {
			for y := 0; y < mirrorRangeY[0]; y++ {
				newX, newRangeX, newY, newRangeY = removeSmudgeRow(note, y, mirrorX, mirrorY)
				if newX != -1 && newX != mirrorX || newY != -1 && newY != mirrorY {
					return newX, newRangeX, newY, newRangeY
				}
			}
		} else {
			for y := mirrorRangeY[1]; y < matrix.LenColumn(note); y++ {
				newX, newRangeX, newY, newRangeY = removeSmudgeRow(note, y, mirrorX, mirrorY)
				if newX != -1 && newX != mirrorX || newY != -1 && newY != mirrorY {
					return newX, newRangeX, newY, newRangeY
				}
			}
		}
	}

	return bruteForceSmudge(note, mirrorX, mirrorRangeX, mirrorY, mirrorRangeY)
}

func bruteForceSmudge(note [][]string, mirrorX int, mirrorRangeX []int, mirrorY int, mirrorRangeY []int) (int, []int, int, []int) {
	newRangeX, newRangeY := []int{-1, -1}, []int{-1, -1}
	newX, newY := -1, -1
	for x := 0; x < matrix.LenRow(note); x++ {
		for y := 0; y < matrix.LenColumn(note); y++ {
			copyNote := matrix.Copy(note)
			if copyNote[x][y] == "#" {
				copyNote[x][y] = "."
			} else {
				copyNote[x][y] = "#"
			}
			newX, newRangeX, newY, newRangeY = findMirrorIndex(copyNote, mirrorX, mirrorY)
			if newX >= 0 && mirrorX != newX || newY >= 0 && mirrorY != newY {
				return newX, newRangeX, newY, newRangeY
			}
		}
	}

	return -1, []int{-1, -1}, -1, []int{-1, -1}
}

func removeSmudgeRow(note [][]string, y int, mirrorX int, mirrorY int) (int, []int, int, []int) {
	row := matrix.Row(y, note)
	newRangeX, newRangeY := []int{-1, -1}, []int{-1, -1}
	newX, newY := -1, -1

	for x, value := range row {
		copyNote := matrix.Copy(note)
		if value == "#" {
			copyNote[x][y] = "."
		} else {
			copyNote[x][y] = "#"
		}

		newX, newRangeX, newY, newRangeY = findMirrorIndex(copyNote, mirrorX, mirrorY)
		if newX >= 0 && mirrorX != newX || newY >= 0 && mirrorY != newY {
			return newX, newRangeX, newY, newRangeY
		}
	}

	return -1, []int{-1, -1}, -1, []int{-1, -1}
}

func removeSmudgeColumn(note [][]string, x int, mirrorX int, mirrorY int) (int, []int, int, []int) {
	column := matrix.Column(x, note)
	newRangeX, newRangeY := []int{-1, -1}, []int{-1, -1}
	newX, newY := -1, -1

	for y, value := range column {
		copyNote := matrix.Copy(note)
		if value == "#" {
			copyNote[x][y] = "."
		} else {
			copyNote[x][y] = "#"
		}

		newX, newRangeX, newY, newRangeY = findMirrorIndex(copyNote, mirrorX, mirrorY)
		if newX >= 0 && mirrorX != newX || newY >= 0 && mirrorY != newY {
			return newX, newRangeX, newY, newRangeY
		}
	}

	return -1, []int{-1, -1}, -1, []int{-1, -1}
}

func findMirrorIndex(note [][]string, skipX int, skipY int) (int, []int, int, []int) {
	rows := matrix.RowStrings(note)
	indexes := findPossibleMirrorIndexes(rows, false)
	valid := false
	rangeX, rangeY := []int{-1, -1}, []int{-1, -1}

	x, y := -1, -1

	if len(indexes) >= 0 {
		for _, index := range indexes {
			if index == skipY {
				continue
			}
			valid, rangeY = validationMirrorIndex(rows, index)
			if valid {
				y = index
				break
			}
			rangeY = []int{-1, -1}
		}
	}

	columns := matrix.ColumnStrings(note)
	indexes = findPossibleMirrorIndexes(columns, true)
	if len(indexes) >= 0 {
		for _, index := range indexes {
			if index == skipX {
				continue
			}
			valid, rangeX = validationMirrorIndex(columns, index)
			if valid {
				x = index
				break
			}
			rangeX = []int{-1, -1}
		}
	}

	return x, rangeX, y, rangeY
}

func validationMirrorIndex(note []string, index int) (bool, []int) {
	lenNote := len(note)
	beforeIndex, afterIndex := -1, -1

	for i := 1; i < lenNote-1; i++ {
		beforeIndex := index - i
		afterIndex := index + i + 1

		if beforeIndex < 0 || afterIndex > lenNote-1 {
			return true, []int{beforeIndex + 1, afterIndex - 1}
		}

		if note[beforeIndex] != note[afterIndex] {
			return false, []int{-1. - 1}
		}
	}
	return true, []int{beforeIndex, afterIndex}
}

func findPossibleMirrorIndexes(feed []string, log bool) []int {
	lenNote := len(feed)
	indexes := []int{}
	for i := 0; i < lenNote-1; i++ {
		current := feed[i]
		next := feed[i+1]

		if current == next {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func parseInput(input []string) [][][]string {
	parsedInput := [][][]string{}
	for _, paragraph := range input {
		m := matrix.Create2D(paragraph)
		parsedInput = append(parsedInput, m)
	}
	return parsedInput
}
