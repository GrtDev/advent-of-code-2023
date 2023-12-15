package utils

import (
	"strconv"
)

func TakeFirstInts(value [][]int) []int {
	if len(value) == 0 {
		return []int{}
	}
	firsts := make([]int, len(value))
	for i, v := range value {
		firsts[i] = v[0]
	}
	return firsts
}

func IntsToString(ints []int) []string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(v)
	}
	return strs
}
