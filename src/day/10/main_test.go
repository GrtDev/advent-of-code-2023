package day10

import (
	"testing"
)

func TestA1(t *testing.T) {
	result, err := RunA("./test-1.txt")
	expected := 4

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestA2(t *testing.T) {
	result, err := RunA("./test-2.txt")
	expected := 8

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
func TestA3(t *testing.T) {
	result, err := RunA("./input.txt")
	expected := 6842

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB(t *testing.T) {
	result, err := RunB("./test-B1.txt")
	expected := 4

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestB2(t *testing.T) {
	result, err := RunB("./test-B2.txt")
	expected := 8

	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestTileCount(t *testing.T) {
	input := []tile{
		tile{
			x:        0,
			tileType: VerticalPipe,
		},
		tile{
			x:        3,
			tileType: VerticalPipe,
		},
	}

	expected := 2
	result := countEnclosedTiles(input)

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestTileCount2(t *testing.T) {
	input := []tile{
		tile{x: 0, tileType: NorthEastPipe},
		tile{x: 1, tileType: HorizontalPipe},
		tile{x: 2, tileType: HorizontalPipe},
		tile{x: 3, tileType: SouthWestPipe},
		tile{x: 8, tileType: SouthEastPipe},
		tile{x: 9, tileType: HorizontalPipe},
		tile{x: 10, tileType: NorthWestPipe},
	}

	expected := 4
	result := countEnclosedTiles(input)

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestTileCount3(t *testing.T) {
	input := []tile{
		tile{x: 1, tileType: VerticalPipe},   // open
		tile{x: 2, tileType: NorthEastPipe},  // count
		tile{x: 3, tileType: HorizontalPipe}, //
		tile{x: 4, tileType: SouthWestPipe},  //
		tile{x: 6, tileType: SouthEastPipe},
		tile{x: 7, tileType: HorizontalPipe},
		tile{x: 8, tileType: NorthWestPipe},
		tile{x: 9, tileType: VerticalPipe},
	}

	expected := 0
	result := countEnclosedTiles(input)

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestTileCount4(t *testing.T) {
	input := []tile{
		tile{x: 1, tileType: VerticalPipe},
		tile{x: 4, tileType: VerticalPipe},
		tile{x: 6, tileType: VerticalPipe},
		tile{x: 9, tileType: VerticalPipe},
	}

	expected := 4
	result := countEnclosedTiles(input)

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}

func TestTileCount5(t *testing.T) {
	input := []tile{
		tile{x: 1, tileType: NorthEastPipe},
		tile{x: 2, tileType: HorizontalPipe},
		tile{x: 3, tileType: HorizontalPipe},
		tile{x: 4, tileType: NorthWestPipe},
		tile{x: 6, tileType: NorthEastPipe},
		tile{x: 7, tileType: HorizontalPipe},
		tile{x: 8, tileType: HorizontalPipe},
		tile{x: 9, tileType: NorthWestPipe},
	}

	expected := 0
	result := countEnclosedTiles(input)

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
func TestTileCount6(t *testing.T) {
	input := []tile{
		tile{x: 1, tileType: VerticalPipe},  //open
		tile{x: 2, tileType: VerticalPipe},  //close
		tile{x: 3, tileType: VerticalPipe},  // open
		tile{x: 4, tileType: NorthEastPipe}, // close
		tile{x: 5, tileType: NorthWestPipe},
		tile{x: 7, tileType: NorthEastPipe},
		tile{x: 8, tileType: SouthWestPipe},
	}

	expected := 1
	result := countEnclosedTiles(input)

	if result != expected {
		t.Fatalf(`expected: %v, to equal: %v`, result, expected)
	}
}
