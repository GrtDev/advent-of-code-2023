package days

import (
	"advent-of-code-2023/utils"
	"testing"
)

func TestDay02A(t *testing.T) {
	input := utils.ReadLines("./inputs/02_test.txt")

	result, err := Day02A(input)
	expected := 8

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestDay02B(t *testing.T) {
	input := utils.ReadLines("./inputs/02_test.txt")

	result, err := Day02B(input)
	expected := 2286

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
