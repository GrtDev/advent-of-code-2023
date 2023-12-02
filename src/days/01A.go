package days

import (
	"advent-of-code-2023/utils"
	"log"
	"regexp"
	"strconv"

	"github.com/thoas/go-funk"
)

func Day01A()(int, error) {
    inputLines := utils.ReadLines("./days/inputs/01A.txt")

    digitRegexp := regexp.MustCompile("\\d")

    values := funk.Map(inputLines, func(line string) int {
        digits := digitRegexp.FindAllString(line, -1)
        lineValue := digits[0] + funk.Last(digits).(string)
        lineNumber, error := strconv.Atoi(lineValue)
        if(error != nil) { log.Fatal(error) }
        return lineNumber
    }).([]int)

    solution := funk.SumInt(values)
    return solution, nil
}
