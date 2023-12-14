package day12

import (
	"testing"
)

func TestA(t *testing.T) {
	result, err := RunA("./test.txt")
	expected := 21

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestA2(t *testing.T) {
	result, err := RunA("./input.txt")
	expected := 7407

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB(t *testing.T) {
	result, err := RunB("./test.txt")
	expected := 525152

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestArrangementCount(t *testing.T) {
	result := countVariants("???.###", []int{1, 1, 3})

	expected := 1

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}
func TestArrangementCount4(t *testing.T) {
	result := countVariants("????", []int{1, 1})
	expected := 3

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}

func TestArrangementCount2(t *testing.T) {
	result := countVariants("?#?#?#?#?#?#?#?", []int{1, 3, 1, 6})
	expected := 1

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}

func TestArrangementCount3(t *testing.T) {
	result := countVariants("?###????????", []int{3, 2, 1})
	expected := 10

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}

func TestArrangementCount6(t *testing.T) {
	result := countVariants("??????#????...#?...", []int{9, 1})
	expected := 3

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}

func TestArrangementCount7(t *testing.T) {
	result := countVariants("???.###????.###????.###????.###????.###", []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3})
	expected := 1

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}

func TestArrangementCountB(t *testing.T) {
	record, damaged := unfold("?###????????", []int{3, 2, 1})
	result := countVariants(record, damaged)
	expected := 506250

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}
func TestArrangementCountB2(t *testing.T) {
	record, damaged := unfold("????.#...#...", []int{4, 1, 1})
	result := countVariants(record, damaged)
	expected := 16

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}
func TestArrangementCountB3(t *testing.T) {
	record, damaged := unfold(".??..??...?##.", []int{1, 1, 3})
	result := countVariants(record, damaged)
	expected := 16384

	if result != expected {
		t.Fatalf(`expected length: %v, to equal length: %v`, result, expected)
	}
}
