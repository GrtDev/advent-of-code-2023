package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

func StringToInts(input string) ([]int, error) {
	values := []int{}
	for _, stringValue := range strings.Split(input, " ") {
		if stringValue == "" {
			continue
		}
		value, err := strconv.Atoi(strings.TrimSpace(stringValue))
		if err != nil {
			fmt.Printf("error %+v\n", err)
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func DropEmptyStrings(input []string) []string {
	return funk.Filter(input, func(x string) bool {
		return x != ""
	}).([]string)
}
