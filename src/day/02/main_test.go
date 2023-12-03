package day02

import (
	"advent-of-code-2023/utils"
	"testing"
)

func TestA(t *testing.T) {
	input := utils.ReadLines("./test.txt")

	result, err := RunA(input)
	expected := 8

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB(t *testing.T) {
	input := utils.ReadLines("./test.txt")

	result, err := RunB(input)
	expected := 2286

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
