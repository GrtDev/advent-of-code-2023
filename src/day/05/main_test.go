package day05

import (
	"testing"
)

func TestA(t *testing.T) {
	result, err := RunA("./test.txt")
	expected := 35

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB(t *testing.T) {
	result, err := RunB("./test.txt")
	expected := 46

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
func TestTransformRange(t *testing.T) {
	source := []indexRange{
		{start: 10, end: 20, length: 11},
	}

	iMap := indexMap{
		sourceId:      "source",
		destinationId: "destination",
		ranges: []indexRangeMap{
			{destination: 0, start: 15, end: 20, length: 6},
			{destination: 0, start: 30, end: 35, length: 6},
		},
	}

	result := calculateIndexRanges(source, iMap)

	expected := []indexRange{
		{start: 0, end: 5, length: 6},
		{start: 10, end: 14, length: 5},
	}

	if len(result) != len(expected) {
		t.Fatalf(`expected: length %v, to equal: %v`, len(result), len(expected))
	}
	if result[0] != expected[0] {
		t.Fatalf(`expected: %v, to equal: %v`, result[0], expected[0])
	}
	if result[1] != expected[1] {
		t.Fatalf(`expected: %v, to equal: %v`, result[0], expected[0])
	}
}

func TestTransformRange2(t *testing.T) {
	source := []indexRange{
		{start: 10, end: 20, length: 11},
		{start: 33, end: 34, length: 2},
	}

	iMap := indexMap{
		sourceId:      "source",
		destinationId: "destination",
		ranges: []indexRangeMap{
			{destination: 0, start: 15, end: 20, length: 6},
			{destination: 0, start: 30, end: 35, length: 6},
		},
	}

	result := calculateIndexRanges(source, iMap)

	expected := []indexRange{
		{start: 0, end: 5, length: 6},
		{start: 10, end: 14, length: 5},
		{start: 3, end: 4, length: 2},
	}

	if len(result) != len(expected) {
		t.Fatalf(`expected: length %v, to equal: %v`, len(result), len(expected))
	}
	if result[0] != expected[0] {
		t.Fatalf(`expected: %v, to equal: %v`, result[0], expected[0])
	}
	if result[1] != expected[1] {
		t.Fatalf(`expected: %v, to equal: %v`, result[0], expected[0])
	}
	if result[2] != expected[2] {
		t.Fatalf(`expected: %v, to equal: %v`, result[0], expected[0])
	}
}

func TestTransformRange3(t *testing.T) {
	source := []indexRange{
		{start: 33, end: 38, length: 6},
	}

	iMap := indexMap{
		sourceId:      "source",
		destinationId: "destination",
		ranges: []indexRangeMap{
			{destination: 0, start: 15, end: 20, length: 6},
			{destination: 0, start: 30, end: 35, length: 6},
		},
	}

	result := calculateIndexRanges(source, iMap)

	expected := []indexRange{
		{start: 3, end: 5, length: 3},
		{start: 36, end: 38, length: 3},
	}

	if len(result) != len(expected) {
		t.Fatalf(`expected: length %v, to equal: %v`, len(result), len(expected))
	}
	if result[0] != expected[0] {
		t.Fatalf(`expected: %v, to equal: %v`, result[0], expected[0])
	}
}
