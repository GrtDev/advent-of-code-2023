package days

import (
	"advent-of-code-2023/utils"
	"regexp"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

type colorCount struct {
	red   int
	blue  int
	green int
}

type game struct {
	id    int
	grabs []colorCount
}

func Day02A(input []string) (int, error) {
	if input == nil {
		input = utils.ReadLines("./days/inputs/02.txt")
	}

	games := parseInput(input)
	gameColorCount := colorCount{
		red:   12,
		green: 13,
		blue:  14,
	}

	sumPossibleGameId := 0
	for _, game := range games {
		if gameIsPossible(game, gameColorCount) {
			sumPossibleGameId += game.id
		}
	}

	return sumPossibleGameId, nil
}

func Day02B(input []string) (int, error) {
	if input == nil {
		input = utils.ReadLines("./days/inputs/02.txt")
	}

	games := parseInput(input)
	sumColorCountPowers := 0
	for _, game := range games {
		maxCount := maxColorCount(game)
		sumColorCountPowers += maxCount.red * maxCount.green * maxCount.blue
	}

	return sumColorCountPowers, nil
}

func parseInput(input []string) []game {
	var games []game
	var gameRegexp = regexp.MustCompile("Game (\\d+):\\s(.*)")

	for _, rawGameData := range input {
		matches := gameRegexp.FindStringSubmatch(rawGameData)

		id, _ := strconv.Atoi(matches[1])
		rawGrabs := strings.Split(matches[2], ";")
		grabs := parseRawGrabs(rawGrabs)
		games = append(games, game{
			id:    id,
			grabs: grabs,
		})
	}

	return games
}

/**
 * Parse raw grabs into a slice of colorCount structs
 * A raw grab is a string such as "1 red; 2 green; 3 blue"
 */
func parseRawGrabs(rawGrabs []string) []colorCount {
	var colorsRegexp = regexp.MustCompile("(\\d+) (red|green|blue)")

	return funk.Map(rawGrabs, func(rawGrab string) colorCount {
		colorData := colorsRegexp.FindAllStringSubmatch(rawGrab, -1)

		grab := colorCount{
			red:   0,
			green: 0,
			blue:  0,
		}

		for _, colorCount := range colorData {
			switch colorCount[2] {
			case "red":
				grab.red, _ = strconv.Atoi(colorCount[1])
			case "green":
				grab.green, _ = strconv.Atoi(colorCount[1])
			case "blue":
				grab.blue, _ = strconv.Atoi(colorCount[1])
			}
		}

		return grab

	}).([]colorCount)
}

func gameIsPossible(game game, gameColorCount colorCount) bool {
	maxGameCount := maxColorCount(game)

	return maxGameCount.red <= gameColorCount.red &&
		maxGameCount.green <= gameColorCount.green &&
		maxGameCount.blue <= gameColorCount.blue
}

func maxColorCount(game game) colorCount {
	maxCount := colorCount{
		red:   0,
		green: 0,
		blue:  0,
	}

	for _, grab := range game.grabs {
		maxCount.red = funk.MaxInt([]int{maxCount.red, grab.red})
		maxCount.green = funk.MaxInt([]int{maxCount.green, grab.green})
		maxCount.blue = funk.MaxInt([]int{maxCount.blue, grab.blue})
	}

	return maxCount
}
