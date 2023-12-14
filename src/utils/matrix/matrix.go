package matrix

import (
	"advent-of-code-2023/utils"
	"strings"
)

func Create2D(input string) [][]string {
	rows := utils.ToLines(input)
	lenCols := len(rows)
	matrix := make([][]string, len(rows[0]))
	for y, line := range rows {
		values := strings.Split(line, "")
		for x, v := range values {
			if y == 0 {
				matrix[x] = make([]string, lenCols)
			}
			matrix[x][y] = v
		}
	}
	return matrix
}

func LenColumn(matrix [][]string) int {
	return len(matrix[0])
}

func LenRow(matrix [][]string) int {
	return len(matrix)
}

func Copy(matrix [][]string) [][]string {
	newMatrix := make([][]string, len(matrix))
	for x, col := range matrix {
		newMatrix[x] = make([]string, len(col))
		copy(newMatrix[x], col)
	}
	return newMatrix
}

func Column(x int, matrix [][]string) []string {
	return matrix[x]
}

func ColumnString(x int, matrix [][]string) string {
	column := Column(x, matrix)
	return strings.Join(column, "")
}

func ColumnStrings(matrix [][]string) []string {
	lenColumns := len(matrix)
	columns := make([]string, lenColumns)
	for x := 0; x < lenColumns; x++ {
		columns[x] = ColumnString(x, matrix)
	}
	return columns
}

func Row(y int, matrix [][]string) []string {
	lenRow := len(matrix)
	row := make([]string, lenRow)
	for x := 0; x < lenRow; x++ {
		row[x] += matrix[x][y]
	}
	return row
}

func RowString(y int, matrix [][]string) string {
	return strings.Join(Row(y, matrix), "")
}

func RowStrings(matrix [][]string) []string {
	lenRows := len(matrix[0])
	rows := make([]string, lenRows)
	for y := 0; y < lenRows; y++ {
		rows[y] = RowString(y, matrix)
	}
	return rows
}

func Print(matrix [][]string) string {
	print := ""
	for y := 0; y < len(matrix[0]); y++ {
		print += strings.Join(Row(y, matrix), " ") + "\n"
	}
	return print
}
