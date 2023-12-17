package matrix

import (
	"advent-of-code-2023/utils"
	"fmt"
	"strings"
)

type Value interface{}

type Matrix[T interface{}] struct {
	Values [][]T
	Width  int
	Height int
}

func Create[T interface{}](rows [][]T) Matrix[T] {
	height := len(rows)
	width := len(rows[0])
	values := make([][]T, width)
	for y, row := range rows {
		for x, v := range row {
			if y == 0 {
				values[x] = make([]T, height)
			}
			values[x][y] = v
		}
	}

	return Matrix[T]{
		Values: values,
		Width:  width,
		Height: height,
	}
}

func CreateFromString(input string) Matrix[string] {
	lines := utils.ToLines(input)
	values := [][]string{}

	for _, line := range lines {
		values = append(values, strings.Split(line, ""))
	}

	return Create(values)
}

func (m *Matrix[T]) Get(x int, y int) T {
	if x < 0 || y < 0 || x >= m.Width || y >= m.Height {
		var result T
		return result
	}
	return m.Values[x][y]
}

func (m *Matrix[T]) Row(y int) []T {
	row := make([]T, m.Width)
	for x := 0; x < m.Width; x++ {
		row[x] = m.Values[x][y]
	}
	return row
}

func (m *Matrix[T]) Column(x int) []T {
	return m.Values[x]
}

func (m *Matrix[T]) Length() int {
	return m.Width * m.Height
}

func (m *Matrix[T]) Copy() Matrix[T] {
	values := make([][]T, m.Width)

	for x, col := range m.Values {
		values[x] = make([]T, len(col))
		copy(values[x], col)
	}
	return Matrix[T]{
		Values: values,
		Width:  m.Width,
		Height: m.Height,
	}
}

func (m *Matrix[T]) Print() string {
	print := "\n"
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			print += fmt.Sprintf("%v", m.Values[x][y])
		}
		print += "\n"
	}
	return print
}

func (m *Matrix[T]) AllValues() []T {
	values := []T{}
	for _, col := range m.Values {
		values = append(values, col...)
	}
	return values
}

func (m *Matrix[T]) RotateRight() Matrix[T] {
	values := make([][]T, m.Height)
	for x := 0; x < m.Height; x++ {
		values[x] = make([]T, m.Width)
		for y := 0; y < m.Width; y++ {
			values[x][y] = m.Values[y][m.Height-x-1]
		}
	}
	return Matrix[T]{
		Values: values,
		Width:  m.Height,
		Height: m.Width,
	}
}

func (m *Matrix[T]) RotateLeft() Matrix[T] {
	values := make([][]T, m.Height)
	for x := 0; x < m.Height; x++ {
		values[x] = make([]T, m.Width)
		for y := 0; y < m.Width; y++ {
			values[x][y] = m.Values[m.Width-y-1][x]
		}
	}
	return Matrix[T]{
		Values: values,
		Width:  m.Height,
		Height: m.Width,
	}
}

func ColumnString(x int, m Matrix[string]) string {
	return strings.Join(m.Column(x), "")
}

func ColumnStrings(m Matrix[string]) []string {
	columns := make([]string, m.Width)
	for x := 0; x < m.Width; x++ {
		columns[x] = ColumnString(x, m)
	}
	return columns
}

func RowString(y int, m Matrix[string]) string {
	return strings.Join(m.Row(y), "")
}

func RowStrings(m Matrix[string]) []string {
	rows := make([]string, m.Height)
	for y := 0; y < m.Height; y++ {
		rows[y] = RowString(y, m)
	}
	return rows
}

func ReplaceAll[T comparable](old T, new T, m Matrix[T]) Matrix[T] {
	copyMatrix := m.Copy()
	for x, col := range m.Values {
		for y, v := range col {
			if v == old {
				m.Values[x][y] = new
			}
		}
	}
	return copyMatrix
}
