package day07

import (
	"testing"
)

func TestA(t *testing.T) {
	result, err := RunA("./test.txt")
	expected := 6440

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB(t *testing.T) {
	result, err := RunB("./test.txt")
	expected := 5905

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
