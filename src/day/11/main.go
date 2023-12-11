package day11

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"

	"github.com/thoas/go-funk"
)

type point struct {
	x int
	y int
}

type galaxy struct {
	id int
	x  int
	y  int
}

func (g galaxy) String() string {
	return fmt.Sprintf("%v (%v, %v)", g.id, g.x, g.y)
}

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadLines(inputFile)
	} else {
		return utils.ReadLines("./day/11/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	input := getInput(inputFile)
	matrix := parseInput(input)
	matrix = markMultipliers(matrix)
	printMatrix(matrix)
	galaxies := getGalaxies(matrix)
	sets := getSets(galaxies)

	total := 0
	for _, set := range sets {
		route := getRoute(set[0], set[1])
		total += calcRouteLength(route, 2, matrix)
	}
	return total, nil
}

func RunB(inputFile string) (int, error) {
	input := getInput(inputFile)
	matrix := parseInput(input)
	matrix = markMultipliers(matrix)
	printMatrix(matrix)
	galaxies := getGalaxies(matrix)
	sets := getSets(galaxies)

	total := 0
	for _, set := range sets {
		route := getRoute(set[0], set[1])
		total += calcRouteLength(route, 1000000, matrix)
	}
	return total, nil
}

func parseInput(input []string) [][]string {
	matrix := make([][]string, len(input))
	for y, line := range input {
		values := strings.Split(line, "")
		matrix[y] = values
	}
	return matrix
}

func markMultipliers(matrix [][]string) [][]string {
	width := len(matrix[0])

	for x := 0; x < width; x++ {
		column := getColumn(x, matrix)
		if !funk.Contains(column, "#") {
			matrix = setColumnValue(x, "X", matrix)
		}
	}

	height := len(matrix)
	for y := 0; y < height; y++ {
		row := matrix[y]
		if !funk.Contains(row, "#") {
			matrix = setRowValue(y, "X", matrix)
		}
	}

	return matrix
}

func getColumn(x int, matrix [][]string) []string {
	column := make([]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		column[i] = matrix[i][x]
	}
	return column
}

func setColumnValue(x int, value string, matrix [][]string) [][]string {
	for y := 0; y < len(matrix); y++ {
		matrix[y][x] = value
	}
	return matrix
}

func setRowValue(y int, value string, matrix [][]string) [][]string {
	for x := 0; x < len(matrix[0]); x++ {
		matrix[y][x] = value
	}
	return matrix
}

func printMatrix(matrix [][]string) {
	for _, line := range matrix {
		fmt.Printf("%v\n", line)
	}
}

func getGalaxies(matrix [][]string) []galaxy {
	galaxies := []galaxy{}
	idCount := 0

	for y, line := range matrix {
		for x, value := range line {
			if value == "#" {
				idCount++
				galaxies = append(galaxies, galaxy{
					id: idCount,
					x:  x,
					y:  y,
				})
			}
		}
	}

	return galaxies
}

func getSets(galaxies []galaxy) [][]galaxy {
	sets := [][]galaxy{}
	lenG := len(galaxies)
	for i := 0; i < lenG; i++ {
		galaxy1 := galaxies[i]

		for j := i + 1; j < lenG; j++ {
			galaxy2 := galaxies[j]
			sets = append(sets, []galaxy{galaxy1, galaxy2})
		}
	}
	return sets
}

func getRoute(a galaxy, b galaxy) []point {

	x := b.x - a.x
	y := b.y - a.y

	route := []point{}
	for i := 0; i < utils.AbsInt(x); i++ {
		if x < 0 {
			route = append(route, point{a.x + i*-1, a.y})
		} else {
			route = append(route, point{a.x + i, a.y})
		}
	}

	for j := 0; j < utils.AbsInt(y); j++ {
		if y < 0 {
			route = append(route, point{a.x + x, a.y + j*-1})
		} else {
			route = append(route, point{a.x + x, a.y + j})
		}
	}

	return route
}

func calcRouteLength(route []point, multiplier int, matrix [][]string) int {
	length := 0
	for _, point := range route {
		if matrix[point.y][point.x] == "X" {
			length = length + 1*multiplier
		} else {
			length++
		}

	}
	return length
}
