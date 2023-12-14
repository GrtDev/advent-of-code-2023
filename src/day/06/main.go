package day06

import (
	"advent-of-code-2023/utils"
	"regexp"
	"strconv"
	"strings"
)

type raceResult struct {
	durationMs int
	distanceMm int
}

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadLines(inputFile)
	} else {
		return utils.ReadLines("./day/06/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	raceResults := parseRaceResultsA(getInput(inputFile))

	total := 0
	for _, race := range raceResults {
		winningDurations := calculateWinningDurations(race)
		if total > 0 {
			total = total * len(winningDurations)
		} else {
			total = len(winningDurations)
		}
	}

	return total, nil
}

func RunB(inputFile string) (int, error) {
	race := parseRaceResultB(getInput(inputFile))

	winningDurations := calculateWinningDurations(race)

	return len(winningDurations), nil

}

func parseRaceResultsA(input []string) []raceResult {
	raceTimes, _ := utils.StringToInts(strings.Split(input[0], ":")[1], " ")
	raceDistances, _ := utils.StringToInts(strings.Split(input[1], ":")[1], " ")

	raceResults := []raceResult{}
	for i, raceTime := range raceTimes {
		raceResults = append(raceResults, raceResult{
			durationMs: raceTime,
			distanceMm: raceDistances[i],
		})
	}
	return raceResults
}

func parseRaceResultB(input []string) raceResult {
	whitespaceRegexp := regexp.MustCompile(`\s+`)
	raceTime, _ := strconv.Atoi(whitespaceRegexp.ReplaceAllString(strings.Split(input[0], ":")[1], ""))
	raceDistance, _ := strconv.Atoi(whitespaceRegexp.ReplaceAllString(strings.Split(input[1], ":")[1], ""))

	return raceResult{
		durationMs: raceTime,
		distanceMm: raceDistance,
	}
}

func calculateWinningDurations(result raceResult) []int {
	winningDurations := []int{}
	for i := 1; i < result.durationMs; i++ {
		speed := i
		distance := speed * (result.durationMs - i)
		if distance > result.distanceMm {
			winningDurations = append(winningDurations, i)
		} else if len(winningDurations) > 0 {
			break
		}
	}
	return winningDurations
}
