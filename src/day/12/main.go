package day12

import (
	"advent-of-code-2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

var hashRegexp = regexp.MustCompile(`#`)
var skippedHashRegexp = regexp.MustCompile(`X`)

func getInput(inputFile string) []string {
	if inputFile != "" {
		return utils.ReadLines(inputFile)
	} else {
		return utils.ReadLines("./day/12/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	input := getInput(inputFile)
	records, damageCount := parseInput(input)

	sumVariants := 0
	for i, record := range records {
		damage := damageCount[i]
		variants := countVariants(record, damage)
		sumVariants += variants
	}

	return sumVariants, nil
}

func RunB(inputFile string) (int, error) {
	input := getInput(inputFile)
	records, damageCount := parseInput(input)

	sumVariants := 0
	for i, record := range records {
		damage := damageCount[i]
		unfoldRecord, unfoldDamaged := unfold(record, damage)
		variants := countVariants(unfoldRecord, unfoldDamaged)
		sumVariants += variants
	}

	return sumVariants, nil
}

func unfold(record string, damaged []int) (string, []int) {
	unfoldRecord := record
	unfoldDamaged := damaged

	for i := 0; i < 4; i++ {
		unfoldRecord += "?" + record
		unfoldDamaged = append(unfoldDamaged, damaged...)
	}

	return unfoldRecord, unfoldDamaged
}

func parseInput(input []string) ([]string, [][]int) {
	records := []string{}
	damaged := [][]int{}

	for _, line := range input {
		inputParts := strings.Split(line, " ")
		records = append(records, inputParts[0])
		damagedCount, _ := utils.StringToInts(inputParts[1], ",")
		damaged = append(damaged, damagedCount)
	}
	return records, damaged
}

func memoizeKey(record string, damaged []int) string {
	return fmt.Sprintf("%s-%v", record, damaged)
}

var countVariantsMap = map[string]int{}

func countVariants(record string, damaged []int) int {
	key := memoizeKey(record, damaged)
	if collectVariantCount, ok := countVariantsMap[key]; ok {
		return collectVariantCount
	}

	if len(damaged) == 1 {
		return countLastVariants(record, damaged)
	}

	variantsCount := 0
	offset := 0
	recordPrefilled, restIndex := fillRecord(record, damaged[:1], 0, offset)
	sumDamaged := funk.SumInt(damaged)
	lenRecord := len(record)

	for restIndex != -1 &&
		(restIndex < lenRecord) &&
		(sumDamaged+offset < lenRecord) &&
		!skippedHashRegexp.MatchString(recordPrefilled) {

		rest := recordPrefilled[restIndex+1:]
		variantsCount += countVariants(rest, damaged[1:])

		offset++
		newPrefill, newRestIndex := fillRecord(record, damaged[:1], 0, offset)
		for newPrefill == recordPrefilled {
			offset++
			newPrefill, newRestIndex = fillRecord(record, damaged[:1], 0, offset)
		}
		recordPrefilled = newPrefill
		restIndex = newRestIndex
	}

	countVariantsMap[key] = variantsCount
	return variantsCount
}

var lastVariantMap = map[string]int{}

func countLastVariants(prefilledRecord string, damaged []int) int {

	key := memoizeKey(prefilledRecord, damaged)
	if lastVariantsCount, ok := lastVariantMap[key]; ok {
		return lastVariantsCount
	}

	variantsCount := 0
	attempt, lastAttempt := "", ""
	restIndex := -1
	runOffset := 0

	for true {
		attempt, restIndex = fillRecord(prefilledRecord, damaged, len(damaged)-1, runOffset)
		runOffset++

		if restIndex == -1 {
			break
		}
		if skippedHashRegexp.MatchString(attempt) || hashRegexp.MatchString(attempt) || attempt == lastAttempt {
			continue // invalid result
		}
		lastAttempt = attempt
		variantsCount++
	}

	lastVariantMap[key] = variantsCount
	return variantsCount
}

func fillRecord(record string, damaged []int, i int, offset int) (string, int) {
	restIndex := -1
	for j := 0; j < len(damaged); j++ {
		if j == i {
			record, restIndex = fillDamagedRecord(record, damaged[j], offset)
		} else {
			record, restIndex = fillDamagedRecord(record, damaged[j], 0)
		}
		if restIndex == -1 {
			return "", -1
		}
	}

	return record, restIndex
}

func fillDamagedRecord(record string, damageCount int, offset int) (string, int) {
	damageNumberFill := strings.Repeat(strconv.Itoa(damageCount), damageCount)
	offsetPattern := fmt.Sprintf(`[^\d]{%d}`, offset)
	unknownPattern := `[.?\d]*?`
	nonDigitPattern := `^|[^\d]{1}`
	pattern := fmt.Sprintf("[?#]{%d}", damageCount)
	restPattern := `[^#].*|$`
	pattern = fmt.Sprintf(`^(%s)(%s)(%s)(%s)(%s)$`, unknownPattern, nonDigitPattern, offsetPattern, pattern, restPattern)

	reg := regexp.MustCompile(pattern)
	g := reg.FindStringSubmatchIndex(record)

	if len(g) != 12 {
		return "", -1
	}

	prefill := strings.ReplaceAll(record[g[2]:g[7]], "?", ".")
	prefill = strings.ReplaceAll(prefill, "#", "X")
	return prefill + damageNumberFill + record[g[10]:g[11]], g[10]
}
