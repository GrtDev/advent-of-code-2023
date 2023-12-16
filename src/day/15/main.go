package dayxx

import (
	"advent-of-code-2023/utils"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var memCalcHash = utils.Memoize(calcHash)
var boxes = make([][]lens, 256)
var stepRegexp = regexp.MustCompile(`^(\w+)([=-])(\d*)$`)

type lens struct {
	focal int
	label string
}

func getInput(inputFile string) string {
	if inputFile != "" {
		return utils.ReadFile(inputFile)
	} else {
		return utils.ReadFile("./day/15/input.txt")
	}
}

func RunA(inputFile string) (int, error) {
	input := strings.Trim(getInput(inputFile), "\n")
	steps := strings.Split(input, ",")

	total := 0
	for _, step := range steps {
		value := memCalcHash.Call(step)
		total += value
	}

	return total, nil
}

func calcHash(step string) int {
	value := 0
	for _, runeValue := range step {
		value += int(runeValue)
		value = (value * 17) % 256
	}
	return value
}

func RunB(inputFile string) (int, error) {
	input := strings.Trim(getInput(inputFile), "\n")
	steps := strings.Split(input, ",")

	for _, step := range steps {
		processLens(step)
	}

	total := 0
	for i, box := range boxes {
		if box == nil {
			continue
		}
		for l, lens := range box {
			total += (i + 1) * (l + 1) * lens.focal
		}
	}

	return total, nil
}

func processLens(step string) {
	matches := stepRegexp.FindAllStringSubmatch(step, -1)
	if len(matches) == 0 {
		log.Fatalf("step %v doesn't match", step)
	}

	label := matches[0][1]
	boxNumber := memCalcHash.Call(label)
	operation := matches[0][2]
	focalLength := matches[0][3]

	box := boxes[boxNumber]
	lensIndex := -1
	if box == nil {
		box = []lens{}
	} else {
		lensIndex = slices.IndexFunc(box, func(l lens) bool {
			return l.label == label
		})
	}

	if operation == "=" {
		focalLengthInt, _ := strconv.Atoi(focalLength)

		lens := lens{
			focal: focalLengthInt,
			label: label,
		}

		if lensIndex != -1 {
			box[lensIndex] = lens
		} else {
			box = append(box, lens)
		}
	} else if lensIndex != -1 {
		box = append(box[:lensIndex], box[lensIndex+1:]...)
	}

	boxes[boxNumber] = box
}
