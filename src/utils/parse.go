package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func StringsToInts(input string) ([]int, error) {
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
