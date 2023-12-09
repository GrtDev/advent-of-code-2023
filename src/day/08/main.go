package day07

import (
	"advent-of-code-2023/utils"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

type Direction int16

const (
	Left Direction = iota
	Right
)

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadLines(inputFile)
	} else {
		return utils.ReadLines("./day/08/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	var input []string = getInput(inputFile)
	sequence := parseSequence(input[0])
	nodes := parseNodes(input[2:])
	steps := calculateSteps("AAA", "ZZZ", nodes, sequence)
	return steps, nil
}

func RunB(inputFile string) (int, error) {
	var input []string = getInput(inputFile)
	sequence := parseSequence(input[0])
	nodes := parseNodes(input[2:])
	startingNumbers := funk.FilterString(utils.Keys(nodes), func(node string) bool {
		return node[2:] == "A"
	})

	steps := funk.Map(startingNumbers, func(node string) int {
		return stepsToZ(node, nodes, sequence)
	}).([]int)

	intersection := findMultiplicationIntersection(steps)

	return intersection, nil
}

func parseSequence(input string) []Direction {
	sequenceStrings := strings.Split(strings.Trim(input, " "), "")

	sequence := []Direction{}

	for _, dirString := range sequenceStrings {
		switch dirString {
		case "L":
			sequence = append(sequence, Left)
		case "R":
			sequence = append(sequence, Right)
		default:
			panic("invalid direction")
		}
	}

	return sequence
}

func parseNodes(input []string) map[string][]string {
	nodeMap := make(map[string][]string)
	nodeRegexp := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

	for _, line := range input {
		matches := nodeRegexp.FindStringSubmatch(line)
		nodeMap[matches[1]] = []string{matches[2], matches[3]}
	}

	return nodeMap
}

func calculateSteps(start string, target string, nodesMap map[string][]string, sequence []Direction) int {
	node := start
	nodePath := nodesMap[node]
	i := 0
	leni := len(sequence)

	for node != target {
		direction := sequence[i%leni]
		node = nodePath[direction]
		nodePath = nodesMap[node]
		i++
	}

	return i
}

func step(node string, nodesMap map[string][]string, dir Direction) string {
	nodePath := nodesMap[node]
	nextNode := nodePath[dir]
	return nextNode
}

func stepsToZ(node string, nodesMap map[string][]string, sequence []Direction) int {
	nodePaths := nodesMap[node]
	i := 0
	leni := len(sequence)

	for true {
		direction := sequence[i%leni]
		node = nodePaths[direction]
		nodePaths = nodesMap[node]

		if node[2:] == "Z" {
			return i + 1
		}

		i++
	}

	return -1
}

func findMultiplicationIntersection(numbers []int) int {
	smallestNumber := funk.MinInt(numbers)
	i := 0
	value := 0
	lenNumbers := len(numbers)

	for true {
		i++
		value = smallestNumber * i

		for y, number := range numbers {
			if value%number != 0 {
				break
			}
			if y == lenNumbers-1 {
				return value
			}
		}

	}

	return -1
}
