package day02

import (
	"testing"
)

func TestA(t *testing.T) {
	result, err := RunA("./test.txt")
	expected := 8

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB(t *testing.T) {
	result, err := RunB("./test.txt")
	expected := 2286

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
