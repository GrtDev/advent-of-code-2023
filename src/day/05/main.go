package day05

import (
	"advent-of-code-2023/utils"
	"log"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
)

type indexRange struct {
	start  int
	end    int
	length int
}

type indexRangeMap struct {
	destination int
	start       int
	end         int
	length      int
}

type indexMap struct {
	sourceId      string
	destinationId string
	ranges        []indexRangeMap
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

	seedRanges := []indexRange{}
	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		length := seeds[i+1]

		seedRanges = append(seedRanges, indexRange{
			start:  seed,
			end:    seed + length - 1,
			length: length,
		})
	}

	processedRanges := transformIndexRange(seedRanges, "seed", "location", maps)

	lowestSeedLocations := funk.Map(processedRanges, func(iRange indexRange) int {
		return iRange.start
	}).([]int)

	return funk.MinInt(lowestSeedLocations), nil

}

func parseSeedsString(input string) []int {
	seedsRegexp := regexp.MustCompile(`seeds:([\d\s]+)`)
	match := seedsRegexp.FindStringSubmatch(input)
	seeds, _ := utils.StringToInts(match[1], " ")
	return seeds
}

func parseMaps(input []string) []indexMap {
	mapRegexp := regexp.MustCompile(`([\w-]+)-to-([\w-]+) map:([\d\s\n]+)`)
	maps := []indexMap{}

	for _, paragraph := range input {
		matches := mapRegexp.FindStringSubmatch(paragraph)
		rangeStrings := utils.DropEmptyStrings(strings.Split(matches[3], "\n"))

		ranges := funk.Map(rangeStrings, func(rangeString string) indexRangeMap {
			rangeNumbers, _ := utils.StringToInts(rangeString, " ")
			return indexRangeMap{
				destination: rangeNumbers[0],
				start:       rangeNumbers[1],
				length:      rangeNumbers[2],
				end:         rangeNumbers[1] + rangeNumbers[2] - 1,
			}
		}).([]indexRangeMap)

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
		if sourceIndex >= indexRange.start && sourceIndex <= indexRange.start+indexRange.length {
			return indexRange.destination + (sourceIndex - indexRange.start)
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

func calculateIndexRanges(input []indexRange, iMap indexMap) []indexRange {
	ranges := []indexRange{}
	for _, iRange := range input {
		ranges = append(ranges, processMaps(iRange, iMap.ranges)...)
	}

	return funk.Filter(ranges, func(iRange indexRange) bool {
		return iRange.length > 0
	}).([]indexRange)
}

func processMaps(input indexRange, rangeMaps []indexRangeMap) []indexRange {
	results := []indexRange{}
	lenMaps := len(rangeMaps)

	rangeMap := rangeMaps[0]
	before, after, processed := transformRange(input, rangeMap)
	results = append(results, processed)

	if before.length > 0 {
		if lenMaps > 1 {
			results = append(results, processMaps(before, rangeMaps[1:])...)
		} else {
			results = append(results, before)
		}
	}
	if after.length > 0 {
		if lenMaps > 1 {
			results = append(results, processMaps(after, rangeMaps[1:])...)
		} else {
			results = append(results, after)
		}
	}
	return results
}

func transformRange(i indexRange, mapRange indexRangeMap) (indexRange, indexRange, indexRange) {

	if i.end < mapRange.start {
		return i, indexRange{}, indexRange{}
	}
	if i.start > mapRange.end {
		return indexRange{}, i, indexRange{}
	}

	iBefore := indexRange{}
	iAfter := indexRange{}

	if i.start < mapRange.start {
		iBefore = indexRange{
			start:  i.start,
			end:    mapRange.start - 1,
			length: mapRange.start - i.start,
		}
		i.start = mapRange.start
	}

	if i.end > mapRange.end {
		iAfter = indexRange{
			start:  mapRange.end + 1,
			end:    i.end,
			length: i.end - mapRange.end,
		}
		i.end = mapRange.end
	}

	i.length = i.end - i.start

	processStart := funk.MaxInt([]int{i.start, mapRange.start})
	processEnd := funk.MinInt([]int{i.end, mapRange.end})
	delta := mapRange.destination - mapRange.start

	iProcessed := indexRange{
		start:  delta + processStart,
		end:    delta + processEnd,
		length: processEnd - processStart + 1,
	}

	return iBefore, iAfter, iProcessed
}

func transformIndexRange(i []indexRange, from string, to string, maps []indexMap) []indexRange {
	iMap := getIndexMap(from, maps)
	i = calculateIndexRanges(i, iMap)

	if iMap.destinationId == to {
		return i
	}

	return transformIndexRange(i, iMap.destinationId, to, maps)
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
