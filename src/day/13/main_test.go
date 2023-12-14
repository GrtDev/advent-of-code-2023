package day13

import (
	"testing"
)

func TestA(t *testing.T) {
	result, err := RunA("./test.txt")
	expected := 405

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestA2(t *testing.T) {
	result, err := RunA("./input.txt")
	expected := 34772

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestRowA(t *testing.T) {
	result, err := RunA("./test-row.txt")
	expected := 400

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestColumnA(t *testing.T) {
	result, err := RunA("./test-column.txt")
	expected := 5

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestRowB(t *testing.T) {
	result, err := RunB("./test-row.txt")
	expected := 100

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestColumnB(t *testing.T) {
	result, err := RunB("./test-column.txt")
	expected := 300

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB(t *testing.T) {
	result, err := RunB("./test.txt")
	expected := 400

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB2(t *testing.T) {
	result, err := RunB("./test-B.txt")
	expected := 1100

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
func TestB3(t *testing.T) {
	result, err := RunB("./test-B-2.txt")
	expected := 1300

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB4(t *testing.T) {
	result, err := RunB("./test-B-3.txt")
	expected := 100

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
