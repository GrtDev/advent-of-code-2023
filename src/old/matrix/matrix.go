package matrix

import (
	"advent-of-code-2023/utils"
	"slices"
	"strconv"
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

func NumColumns(matrix [][]string) int {
	return LenRow(matrix)
}

func NumRows(matrix [][]string) int {
	return LenColumn(matrix)
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
	print := "\n"
	for y := 0; y < len(matrix[0]); y++ {
		print += strings.Join(Row(y, matrix), " ") + "\n"
	}
	return print
}

func ReplaceString(old string, new string, matrix [][]string) [][]string {
	copyMatrix := Copy(matrix)
	for x, col := range matrix {
		for y, v := range col {
			if v == old {
				copyMatrix[x][y] = new
			}
		}
	}
	return copyMatrix
}

func RotateRight(matrix [][]string) [][]string {
	lenRow := LenRow(matrix)
	lenCol := LenColumn(matrix)
	rotated := make([][]string, lenCol)
	for x := 0; x < lenCol; x++ {
		rotated[x] = make([]string, lenRow)
		for y := 0; y < lenRow; y++ {
			rotated[x][y] = matrix[y][lenCol-x-1]
		}
	}
	return rotated
}

func RotateLeft(matrix [][]string) [][]string {
	lenRow := LenRow(matrix)
	lenCol := LenColumn(matrix)
	rotated := make([][]string, lenCol)
	for x := 0; x < lenCol; x++ {
		rotated[x] = make([]string, lenRow)
		for y := 0; y < lenRow; y++ {
			rotated[x][y] = matrix[lenRow-y-1][x]
		}
	}
	return rotated
}

type Test struct {
	x int
}

type IndexMatrix struct {
	Indexes [][]int
	Width   int
	Height  int
}

func RotateIndexLeft(iMatrix IndexMatrix) IndexMatrix {
	rotated := IndexMatrix{
		Indexes: make([][]int, iMatrix.Height),
		Width:   iMatrix.Height,
		Height:  iMatrix.Width,
	}
	imat := rotated.Indexes
	for x := iMatrix.Width - 1; x >= 0; x-- {
		col := iMatrix.Indexes[x]
		for _, y := range col {
			if imat[y] == nil {
				imat[y] = []int{}
			}
			imat[y] = append(imat[y], iMatrix.Width-x-1)
		}
	}
	return rotated
}

func RotateIndexRight(iMatrix IndexMatrix) IndexMatrix {
	rotated := IndexMatrix{
		Indexes: make([][]int, iMatrix.Height),
		Width:   iMatrix.Height,
		Height:  iMatrix.Width,
	}
	imat := rotated.Indexes
	for x := 0; x < iMatrix.Width; x++ {
		col := iMatrix.Indexes[x]
		for _, y := range col {

			nX := iMatrix.Height - y - 1
			if imat[nX] == nil {
				imat[nX] = []int{}
			}
			imat[nX] = append(imat[nX], x)
		}
	}
	return rotated
}

func PrintIndexMatrix(matrix IndexMatrix) string {
	print := make([][]string, matrix.Width)
	m := matrix.Indexes
	for x := 0; x < matrix.Width; x++ {
		print[x] = make([]string, matrix.Height)
		for y := 0; y < matrix.Height; y++ {
			if slices.Contains(m[x], y) {
				print[x][y] = strconv.Itoa(y)
			} else {
				print[x][y] = "."
			}
		}
	}
	return Print(print)
}

func Values(matrix [][]string) []string {
	values := []string{}
	for _, col := range matrix {
		values = append(values, col...)
	}
	return values
}

func ValuesIndexes(matrix IndexMatrix) []int {
	values := []int{}
	for _, col := range matrix.Indexes {
		values = append(values, col...)
	}
	return values
}
