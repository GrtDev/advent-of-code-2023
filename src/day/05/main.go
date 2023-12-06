package day05

import (
	"advent-of-code-2023/utils"
	"log"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

type indexRange struct {
	destination int
	source      int
	length      int
}

type indexMap struct {
	sourceId      string
	destinationId string
	ranges        []indexRange
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

	seeds := parseSeedsString(input[0])
	maps := parseMaps(input[1:])

	seedLocations := funk.Map(seeds, func(seed int) int {
		return transformIndex(seed, "seed", "location", maps)
	}).([]int)

	return funk.MinInt(seedLocations), nil
}

func RunB(inputFile string) (int, error) {
	var input []string = getInput(inputFile)

	seeds := parseSeedsString(input[0])
	maps := parseMaps(input[1:])

	seedsAdjusted := []int{}
	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		seedCount := seeds[i+1]

		for j := 0; j < seedCount; j++ {
			seedsAdjusted = append(seedsAdjusted, seed+j)
		}
	}

	seedLocations := funk.Map(seedsAdjusted, func(seed int) int {
		return transformIndex(seed, "seed", "location", maps)
	}).([]int)

	return funk.MinInt(seedLocations), nil

}

func parseSeedsString(input string) []int {
	seedsRegexp := regexp.MustCompile(`seeds:([\d\s]+)`)
	match := seedsRegexp.FindStringSubmatch(input)
	seeds, _ := utils.StringToInts(match[1])
	return seeds
}

func parseMaps(input []string) []indexMap {
	mapRegexp := regexp.MustCompile(`([\w-]+)-to-([\w-]+) map:([\d\s\n]+)`)
	maps := []indexMap{}

	for _, paragraph := range input {
		matches := mapRegexp.FindStringSubmatch(paragraph)
		rangeStrings := utils.DropEmptyStrings(strings.Split(matches[3], "\n"))

		ranges := funk.Map(rangeStrings, func(rangeString string) indexRange {
			rangeNumbers, _ := utils.StringToInts(rangeString)
			return indexRange{
				destination: rangeNumbers[0],
				source:      rangeNumbers[1],
				length:      rangeNumbers[2],
			}
		}).([]indexRange)

		maps = append(maps, indexMap{
			sourceId:      matches[1],
			destinationId: matches[2],
			ranges:        ranges,
		})
	}

	return maps
}

func calculateIndex(iMap indexMap, sourceIndex int) int {
	for _, indexRange := range iMap.ranges {
		if sourceIndex >= indexRange.source && sourceIndex <= indexRange.source+indexRange.length {
			return indexRange.destination + (sourceIndex - indexRange.source)
		}
	}
	return sourceIndex
}

func transformIndex(index int, from string, to string, maps []indexMap) int {
	iMap := getIndexMap(from, maps)
	index = calculateIndex(iMap, index)
	if iMap.destinationId == to {
		return index
	}
	return transformIndex(index, iMap.destinationId, to, maps)
}

func getIndexMap(sourceId string, maps []indexMap) indexMap {
	for _, iMap := range maps {
		if iMap.sourceId == sourceId {
			return iMap
		}
	}

	log.Fatalf("could not find map with sourceId %v", sourceId)
	return indexMap{}
}
